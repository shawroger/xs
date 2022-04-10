package config

import (
	"github.com/gookit/config/v2"
)

// ParseFile
//
// 解析文件，获得配置结构体
func ParseFile(filename string) (*Config, error) {

	c := DefaultConfig()

	// 加载配置文件
	err := config.LoadFiles(filename)
	if err != nil {
		return nil, err
	}

	// KEY = ""，绑定所有数据
	err = config.BindStruct("", &c)
	if err != nil {
		return nil, err
	}

	return c, nil

}
