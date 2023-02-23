package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mengdj/goctl-rest-client/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

var (
	version = "v0.0.5"
	author  = []*cli.Author{
		&cli.Author{
			Name:  "mengdj",
			Email: "mengdj@outlook.com",
		},
	}
	commands = []*cli.Command{
		{
			Name:  "rest-client",
			Usage: "generates rest-client factory",
			Action: func(context *cli.Context) error {
				plugin, err := plugin.NewPlugin()
				if nil != err {
					return err
				}
				return generate.Do(plugin)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "package",
					Usage: "the package of rest-client",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Authors = author
	app.Usage = "a plugin of goctl to generate rest-client"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-rest-client: %+v\n", err)
	}
}
