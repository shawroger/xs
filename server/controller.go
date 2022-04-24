package server

import (
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/utils"
	"gitee.com/feimos/xs/xlsx"
	"github.com/gin-gonic/gin"
)

// Context
//
// gin.Context 的同名
type Context = gin.Context

// UseController
//
// 通过加载 *controller.Controller 创建服务器
//
// currentFc 为当前加载的文件配置
func (h *Handler) UseController(cc *controller.Controller, currentFc *config.FileConfig) {

	for index, sheet := range *cc.Object.Sheets {

		var (
			url    string
			prefix string
			suffix string
		)

		if currentFc.Prefix != "" {
			prefix = utils.GenStandardPath(currentFc.Prefix)
			suffix = utils.GenStandardPath(sheet.Name)
			url = utils.JoinStandardPath(prefix, suffix)
		} else {
			url = utils.JoinStandardPath(cc.RawFilePath, sheet.Name)
		}

		// 自动将第一个 sheet 作为 index
		if h.Config.SheetIndex && index == 0 {
			h.debug("sheet %v will be used as index", sheet.Name)
			HandlerCreateService(h, prefix, cc.File, &sheet)

		}

		// 筛选 sheet 并且建立服务
		if currentFc.Sheets == nil || len(currentFc.Sheets) == 0 || utils.Exist(currentFc.Sheets, sheet.Name) {
			HandlerCreateService(h, url, cc.File, &sheet)
		}
	}
}

// HandlerCreateService
//
// 创建服务
func HandlerCreateService(h *Handler, url string, f *xlsx.File, sheet *xlsx.Sheet) {
	// 建立 GET 服务
	h.CreateServiceGet(url, sheet)

	// 建立 POST 服务
	h.CreateServicePost(url, f, sheet)

	// 建立 DELETE 服务
	h.CreateServiceDelete(url, f, sheet)
}
