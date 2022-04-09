package main

import (
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/server"
	"os"
	"path"
)

func main() {

	dir := "testData/xlsx"
	svr := server.New()
	svr.Debug = true
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := path.Join(dir, file.Name())
			c, err := controller.OpenFile(filename)
			if err != nil {
				panic(err)
			}
			svr.UseController(c)
		}
	}

	svr.Run()

}
