package main

import (
	"gitee.com/feimos/xs/server"
)

func main() {

	svr := server.New()
	err := svr.LoadConfigFile("./testdata/.xs.json")
	if err != nil {
		panic(err)
	}

	err = svr.Run()
	if err != nil {
		panic(err)
	}

}
