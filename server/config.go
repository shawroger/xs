package server

import (
	"fmt"
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/datatype"
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

	h.CreateServiceIndex()

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

// BindAppInfo
//
// 从 cmd 中绑定 AppInfo
func (h *Handler) BindAppInfo(a *datatype.AppInfo) *Handler {
	h.AppInfo = a
	return h
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
