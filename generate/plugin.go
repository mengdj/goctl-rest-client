package generate

import (
	_ "embed"
	"errors"
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util"
)

var (
	//go:embed template/client.tpl
	clientTpl string
)

func Do(plugin *plugin.Plugin, context *cli.Context) error {
	client := &Client{
		Destination: context.String("destination"),
		File:        context.String("file"),
		Version:     Version,
		Package:     context.String("package"),
	}
	for _, tt := range plugin.Api.Types {
		//just struct
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
				HandlerDoc:   route.HandlerDoc,
				Doc:          route.Doc,
				Method:       route.Method,
				Path:         group.GetAnnotation("prefix") + route.Path,
				RequestName:  route.RequestTypeName(),
				ResponseName: route.ResponseTypeName(),
				Comment:      route.Docs,
				Text:         strings.ReplaceAll(route.AtDoc.Text, `"`, ``),
			})
		}
	}
	dir := path.Join(plugin.Dir, client.Package)
	if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); nil != err {
		return err
	}
	return util.With("plugin").Parse(clientTpl).GoFmt(true).SaveTo(client, path.Join(dir, client.File), true)
}
