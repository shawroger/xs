package config

func DefaultConfig() *Config {
	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    nil,
	}
}

// NewConfigFromFiles
//
// 从已有文件配置创建 config
func NewConfigFromFiles(files []FileConfig) *Config {
	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    files,
	}
}

// NewConfigFromSingleFile
//
// 从单一文件路径创建 config
func NewConfigFromSingleFile(filename string) *Config {
	files := []FileConfig{CreateDefaultFromFile(filename)}
	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    files,
	}
}
