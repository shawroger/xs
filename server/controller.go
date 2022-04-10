package server

import (
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/utils"
	"github.com/gin-gonic/gin"
)

// Context
//
// gin.Context 的同名
type Context = gin.Context

// UseController
//
// 通过加载 *controller.Controller 创建服务器
func (h *Handler) UseController(cc *controller.Controller, fc *config.FileConfig) {

	for _, sheet := range *cc.Object.Sheets {
		var url string
		if fc.Prefix != "" {
			prefix := utils.GenStandardPath(fc.Prefix)
			suffix := utils.GenStandardPath(sheet.Name)
			url = utils.JoinStandardPath(prefix, suffix)
		} else {
			url = utils.JoinStandardPath(cc.RawFilePath, sheet.Name)
		}

		// 筛选 sheet 并且建立服务
		if fc.Sheets == nil || len(fc.Sheets) == 0 || utils.Exist(fc.Sheets, sheet.Name) {
			h.CreateServiceGet(url, &sheet)
		}
	}
}
