package cmd

import (
	"github.com/urfave/cli/v2"
)

// AppInfo
//
// 配置 cmd 信息
type AppInfo struct {
	Version string
}

// Flags
//
// cmd 的 flag 信息
type Flags struct {
	file       string
	port       int
	sheetIndex bool
}

func AppInit(a *AppInfo) *cli.App {

	app := &cli.App{}

	app.Name = "xs"
	app.Version = a.Version
	app.Usage = "A easy way to run a server by using xlsx files"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Required: true,
			Aliases:  []string{"f"},
			Usage:    "define the config file",
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
		sheetIndex := c.Bool("sheetIndex")

		return serverHandler(Flags{file, port, sheetIndex})
	}

	return app
}
