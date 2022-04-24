package cmd

import "gitee.com/feimos/xs/server"

func serverHandler(flags Flags) error {
	svr := server.New()
	err := svr.LoadConfigFile(flags.file)
	if err != nil {
		return err
	}

	// 重新设置 port
	if flags.port != 0 {
		svr.Config.Port = flags.port
	}

	if flags.sheetIndex {
		svr.Config.SheetIndex = flags.sheetIndex
	}

	return svr.Run()
}
