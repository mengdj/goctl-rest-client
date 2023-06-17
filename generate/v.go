package generate

import "github.com/zeromicro/go-zero/tools/goctl/api/spec"

// //go:generate fieldalignment -fix v.go
//
//go:generate easyjson  v.go
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
		Destination  string
		File         string
		Version      string
		Package      string
		GroupPackage string
		Route        []Route
		Type         []spec.DefineStruct
		Pkg          string
		Date         string
		Mode         string
	}
	//easyjson:json
	JSONListResult struct {
		Path      string `json:"Path"`
		Main      bool   `json:"Main"`
		Dir       string `json:"Dir"`
		GoMod     string `json:"GoMod"`
		GoVersion string `json:"GoVersion"`
	}
)

const Version = "v0.1.0"
