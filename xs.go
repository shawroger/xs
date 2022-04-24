package main

import (
	"gitee.com/feimos/xs/cmd"
	"gitee.com/feimos/xs/datatype"
	"log"
	"os"
)

var appInfo = datatype.AppInfo{
	Version: "0.0.1",
}

func main() {

	app := cmd.AppInit(&appInfo)
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
