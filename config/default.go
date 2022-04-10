package config

func DefaultConfig() *Config {
	return &Config{
		Port:     8080,
		Debug:    true,
		GinDebug: false,
		Files:    nil,
	}
}
