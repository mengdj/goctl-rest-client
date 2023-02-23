package generate

import (
	_ "embed"
	"errors"
	"os"
	"path"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

var (
	//go:embed template/client.tpl
	clientTpl string
)

func Do(plugin *plugin.Plugin) error {
	client := new(Client)
	for _, tt := range plugin.Api.Types {
		if target, ok := tt.(spec.DefineStruct); ok {
			client.Type = append(client.Type, target)
		} else {
			return errors.New("can't support type")
		}

	}
	for _, group := range plugin.Api.Service.Groups {
		for _, route := range group.Routes {
			client.Route = append(client.Route, Route{
				Handler:      route.Handler,
				Method:       route.Method,
				Path:         group.GetAnnotation("prefix") + route.Path,
				RequestName:  route.RequestTypeName(),
				ResponseName: route.ResponseTypeName(),
				Comment:      route.Docs,
				Text:         route.AtDoc.Text,
			})
		}
	}
	dir := path.Join(plugin.Dir, "client")
	if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); nil != err {
		return err
	}
	return util.With("plugin").Parse(clientTpl).GoFmt(true).SaveTo(client, path.Join(dir, "client.go"), true)
}
