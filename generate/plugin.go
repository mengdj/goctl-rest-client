package generate

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/core/mr"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"golang.org/x/mod/modfile"
)

var (
	//go:embed template/clients.tpl
	clientsTpl string
	//go:embed template/types.tpl
	typesTpl string
	goMod    = "go.mod"
	notFound = errors.New("not found go.mod")
)

func foundPkgFromCommand(ctx context.Context, path string) (string /*dir*/, string /*mod*/, error) {
	var (
		cmd = []*exec.Cmd{
			//set env
			exec.CommandContext(ctx, "go", "env", "-w", "GO111MODULE=auto"),
			exec.CommandContext(ctx, "go", "list", "-json", "-m"),
		}
		body   []byte
		err    error
		result *JSONListResult = nil
	)
	for _, c := range cmd {
		c.Dir = filepath.Dir(path)
		if body, err = c.CombinedOutput(); err != nil {
			return "", "", errors.Wrapf(err, string(body))
		}
	}
	//parse
	result = &JSONListResult{}
	if err = result.UnmarshalJSON(body); nil != err {
		return "", "", err
	}
	return result.Dir, result.Path, nil
}

func foundPkg(dir string) (string /*dir*/, string /*mod*/, error) {
	pkg := ""
	entries, err := os.ReadDir(dir)
	if nil != err {
		return "", "", err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		//compare
		if entry.Name() == goMod {
			//read go.mod
			fb, fe := os.ReadFile(path.Join(dir, goMod))
			if nil != fe {
				return "", "", fe
			}
			md, me := modfile.Parse(goMod, fb, nil)
			if nil != me {
				return "", "", me
			}
			if nil == md.Module {
				return "", "", notFound
			}
			pkg = md.Module.Mod.Path
			break
		}
	}
	if "" == pkg {
		//top
		if prevDir := filepath.Dir(dir); prevDir != dir {
			return foundPkg(prevDir)
		} else {
			return "", "", notFound
		}
	}
	return dir, pkg, nil
}

func Do(plugin *plugin.Plugin, context *cli.Context) error {
	var (
		client = &Client{
			Destination: context.String("destination"),
			File:        context.String("file"),
			Version:     Version,
			Package:     context.String("package"),
			Date:        time.Now().String(),
		}
		groupSize = len(plugin.Api.Service.Groups)
		typeSize  = len(plugin.Api.Types)
		module    string
		moduleDir string
		err       error
	)
	//go list -json -m
	if moduleDir, module, err = foundPkgFromCommand(context.Context, plugin.ApiFilePath); nil != err {
		//found go.mod
		if moduleDir, module, err = foundPkg(plugin.Dir); nil != err {
			fmt.Printf("found error:%s", err.Error())
			return err
		}
	}
	//parse package
	// windows=\\ unix=/
	client.Pkg = module + path.Join(strings.ReplaceAll(strings.ReplaceAll(plugin.Dir, moduleDir, ""), "\\", "/"), client.Package)
	return mr.Finish(func() error {
		//build types
		for i := 0; i < typeSize; i++ {
			if target, ok := plugin.Api.Types[i].(spec.DefineStruct); ok {
				client.Type = append(client.Type, target)
			}
		}
		dir := path.Join(plugin.Dir, client.Package)
		if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); nil != err {
			return err
		}
		//build types
		return With("types").Parse(typesTpl).GoFmt(true).SaveTo(client, path.Join(dir, "types.go"), true)
	}, func() error {
		//build groups
		for i := 0; i < groupSize; i++ {
			group := plugin.Api.Service.Groups[i]
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
			client.GroupPackage = strings.ReplaceAll(group.GetAnnotation("group"), `"`, ``)
			dir := path.Join(plugin.Dir, client.Package, client.GroupPackage)
			err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
			if nil != err {
				return err
			}
			if err = With(client.GroupPackage).Parse(clientsTpl).GoFmt(true).SaveTo(client, path.Join(dir, client.File), true); nil != err {
				return err
			}
			client.Route = client.Route[0:0]
		}
		return nil
	})
}
