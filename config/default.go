package config

func DefaultConfig() *Config {
	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    nil,
	}
}

// NewConfigFromSingleFile
//
// 从单一文件路径创建 config
func NewConfigFromSingleFile(filename ...string) *Config {
	files := make([]FileConfig, len(filename))

	for i, name := range filename {
		files[i] = CreateDefaultFromFilename(name)
	}

	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    files,
	}
}
