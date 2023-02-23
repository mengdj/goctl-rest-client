package generate

import "github.com/zeromicro/go-zero/tools/goctl/api/spec"

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
		Comment      []string
		Text         string
	}

	Client struct {
		Route []Route
		Type  []spec.DefineStruct
	}
)
