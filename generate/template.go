// Package generate
// @file:template.go
// @description:
// @date: 02/25/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package generate

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	goformat "go/format"
	"os"
	"text/template"
	"unicode"
)

const regularPerm = 0o666

// DefaultTemplate is a tool to provides the text/template operations
type DefaultTemplate struct {
	name  string
	text  string
	goFmt bool
}

// With returns an instance of DefaultTemplate
func With(name string) *DefaultTemplate {
	return &DefaultTemplate{
		name: name,
	}
}

// Parse accepts a source template and returns DefaultTemplate
func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	t.text = text
	return t
}

// GoFmt sets the value to goFmt and marks the generated codes will be formatted or not
func (t *DefaultTemplate) GoFmt(format bool) *DefaultTemplate {
	t.goFmt = format
	return t
}

// SaveTo writes the codes to the target path
func (t *DefaultTemplate) SaveTo(data interface{}, path string, forceUpdate bool) error {
	var (
		output *bytes.Buffer
		err    error
	)
	if pathx.FileExists(path) && !forceUpdate {
		return nil
	}
	if output, err = t.Execute(data); err != nil {
		return err
	}
	return os.WriteFile(path, output.Bytes(), regularPerm)
}

// Execute returns the codes after the template executed
func (t *DefaultTemplate) Execute(data interface{}) (*bytes.Buffer, error) {
	var (
		tem          *template.Template
		err          error
		buf          *bytes.Buffer
		formatOutput []byte
	)
	if tem, err = template.New(t.name).Funcs(template.FuncMap{
		"isPublic": func(str string) bool { return unicode.IsUpper(rune(str[0])) },
	}).Parse(t.text); err != nil {
		return nil, errors.Wrapf(err, "template parse error:", t.text)
	}
	buf = new(bytes.Buffer)
	if err = tem.Execute(buf, data); err != nil {
		return nil, errors.Wrapf(err, "template execute error:", t.text)
	}
	if !t.goFmt {
		return buf, nil
	}
	if formatOutput, err = goformat.Source(buf.Bytes()); err != nil {
		return nil, errors.Wrapf(err, "go format error:", buf.String())
	}
	buf.Reset()
	buf.Write(formatOutput)
	return buf, nil
}
