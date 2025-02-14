package cmd

import (
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/server"
	"gitee.com/feimos/xs/utils"
)

type ErrMissingDataSource struct {
}

func (s ErrMissingDataSource) Error() string {
	return "Missing flag which provides the data source\nconsider: using either \"-f\" or \"-c\""
}

func serverHandler(flags Flags, a *datatype.AppInfo) error {
	// 从配置文件
	if flags.conf != "" {
		return startServerFromConf(flags, a)
	}

	// 直接读取 xlsx 文件
	if flags.files != "" {
		return startServerFromFile(flags, a)
	}

	return ErrMissingDataSource{}
}

func startServerFromConf(flags Flags, a *datatype.AppInfo) error {

	c, err := config.ParseFile(flags.conf)

	if err != nil {
		return err
	}

	// 重新设置 port
	if flags.port != 0 {
		c.Port = flags.port
	}

	svr := server.New().BindAppInfo(a)
	err = svr.UseConfig(c)

	if err != nil {
		return err
	}

	return svr.Run()
}

func startServerFromFile(flags Flags, a *datatype.AppInfo) error {

	files, err := utils.ParseCmdFilesFlag(flags.files)

	if err != nil {
		return err
	}

	c := config.NewConfigFromSingleFile(files...)
	// 重新设置 port
	if flags.port != 0 {
		c.Port = flags.port
	}

	svr := server.New().BindAppInfo(a)
	err = svr.UseConfig(c)

	if err != nil {
		return err
	}

	return svr.Run()
}
