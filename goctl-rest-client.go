package main

import (
	"fmt"
	"github.com/mengdj/goctl-rest-client/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"os"
	"runtime"
)

var (
	author = []*cli.Author{
		&cli.Author{
			Name:  "mengdj",
			Email: "mengdj@outlook.com",
		},
	}
	commands = []*cli.Command{
		{
			Name:  "rest-client",
			Usage: "generates rest-client",
			Action: func(context *cli.Context) error {
				plugin, err := plugin.NewPlugin()
				if nil != err {
					return err
				}
				return generate.Do(plugin, context)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "package",
					Usage:    "the package of rest-client",
					Value:    "client",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "destination",
					Usage:    "destination address,for example www.baidu.com or service",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "file",
					Usage:    "target file",
					Required: false,
					Value:    "client.go",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Authors = author
	app.Usage = "a plugin of goctl to generate rest-client"
	app.Version = fmt.Sprintf("%s %s/%s", generate.Version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-rest-client: %+v\n", err)
	}
}
