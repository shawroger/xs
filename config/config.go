package config

import "gitee.com/feimos/xs/utils"

// Config
//
// 配置项 结构体
type Config struct {
	Port     int
	Debug    bool
	GinDebug bool
	Files    []FileConfig
}

// FileConfig
//
// 文件配置项
type FileConfig struct {
	// 文件路径
	File string
	// 使用的前缀
	Prefix string
	// 筛选工作表
	//
	// 为空，则是为全部开启
	Sheets []string
}

// CreateDefaultFromFilename
//
// 从单文件路径生成默认 FileConfig
func CreateDefaultFromFilename(file string) FileConfig {
	prefix := utils.GetPureName(file)
	return FileConfig{
		File:   file,
		Prefix: prefix,
	}
}
