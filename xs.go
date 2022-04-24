package main

import (
	"gitee.com/feimos/xs/cmd"
	"log"
	"os"
)

var appInfo = cmd.AppInfo{
	Version: "0.0.1",
}

func main() {

	app := cmd.AppInit(&appInfo)
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
