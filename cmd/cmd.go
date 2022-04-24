package cmd

import (
	"gitee.com/feimos/xs/datatype"
	"github.com/urfave/cli/v2"
)

// Flags
//
// cmd 的 flag 信息
type Flags struct {
	files string
	conf  string
	port  int
}

func AppInit(a *datatype.AppInfo) *cli.App {

	app := &cli.App{}

	app.Name = "xs"
	app.Version = a.Version
	app.Usage = "A easy way to run a server by using xlsx files"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "conf",
			Aliases: []string{"c"},
			Usage:   "define the config file",
		}, &cli.StringFlag{
			Name:    "files",
			Aliases: []string{"f"},
			Usage:   "define the xlsx file",
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "define the port of server",
		},
	}

	app.Action = func(c *cli.Context) error {
		port := c.Int("port")
		conf := c.String("conf")
		files := c.String("files")

		return serverHandler(Flags{files, conf, port}, a)
	}

	return app
}
