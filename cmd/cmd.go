package cmd

import (
	"gitee.com/feimos/xs/datatype"
	"github.com/urfave/cli/v2"
)

// Flags
//
// cmd 的 flag 信息
type Flags struct {
	file       string
	conf       string
	port       int
	sheetIndex bool
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
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "define the xlsx file",
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "define the port of server",
		},
		&cli.BoolFlag{
			Name:    "sheetIndex",
			Aliases: []string{"i"},
			Usage:   "Sheet1 will be used as index",
		},
	}

	app.Action = func(c *cli.Context) error {
		port := c.Int("port")
		file := c.String("file")
		conf := c.String("conf")
		sheetIndex := c.Bool("sheetIndex")

		return serverHandler(Flags{file, conf, port, sheetIndex}, a)
	}

	return app
}
