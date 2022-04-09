package server

import (
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/utils"
)

// UseController
//
// 通过加载 *controller.Controller 创建服务器
func (h *Handler) UseController(cc *controller.Controller) {

	for _, sheet := range *cc.Object.Sheets {
		url := utils.GenStandardPath(cc.RawFilePath) + "/" + utils.GenStandardPath(sheet.Name)

		h.SetRequestGet(url, &sheet)
	}
}
