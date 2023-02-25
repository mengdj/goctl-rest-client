package generate

import "github.com/zeromicro/go-zero/tools/goctl/api/spec"

//go:generate fieldalignment -fix v.go
type (
	Method struct {
		Name     string
		Request  string
		Response string
	}

	Type struct {
		Name     string
		Document []string
		Comment  []string
	}

	Route struct {
		Handler      string
		Method       string
		Path         string
		RequestName  string
		ResponseName string
		Text         string
		Comment      []string
		HandlerDoc   []string
		Doc          spec.Doc
	}

	Client struct {
		Destination string
		File        string
		Version     string
		Package     string
		Route       []Route
		Type        []spec.DefineStruct
	}
)

const Version = "v0.0.2"
