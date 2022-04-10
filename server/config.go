package server

import (
	"fmt"
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/controller"
)

// UseConfig
//
// 加载配置
func (h *Handler) UseConfig(c *config.Config) error {
	h.Config = c

	// 建立 controller
	for _, fc := range h.Config.Files {
		filename := fc.File
		cc, err := controller.OpenFile(filename)
		if err != nil {
			return err
		}

		h.UseController(cc, &fc)
	}

	return nil
}

// LoadConfigFile
//
// 从配置文件中加载配置
func (h *Handler) LoadConfigFile(filename string) error {

	c, err := config.ParseFile(filename)
	if err != nil {
		return err
	}

	err = h.UseConfig(c)
	if err != nil {
		return err
	}
	return nil
}

// ErrMissingHandlerConfig
//
// 自定义错误：未定义配置
type ErrMissingHandlerConfig struct {
}

func (c ErrMissingHandlerConfig) Error() string {
	return fmt.Sprintf("Missing config struct in server handler\nconsider: use \"UseConfig\" or \"LoadConfigFile\" methods")
}

// GetConfig
//
// 读取配置
func (h *Handler) GetConfig() (*config.Config, error) {
	if h.Config == nil {
		return nil, ErrMissingHandlerConfig{}
	}

	return h.Config, nil
}
