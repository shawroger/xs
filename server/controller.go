package server

import (
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/controller"
	"gitee.com/feimos/xs/utils"
	"gitee.com/feimos/xs/xlsx"
	"github.com/gin-gonic/gin"
	"strings"
)

// Context
//
// gin.Context 的同名
type Context = gin.Context

// 如果字符串有空格则替换为其他字符，并给出警告
func checkPathWithSpace(h *Handler, rawPath, replace string) string {
	if strings.Contains(rawPath, " ") {
		newPath := strings.ReplaceAll(rawPath, " ", replace)
		h.debug("Warning: \"%v\" has space , it will be rename as \"%v\"", rawPath, newPath)
		return newPath
	}

	return rawPath
}

// UseController
//
// 通过加载 *controller.Controller 创建服务器
//
// currentFc 为当前加载的文件配置
func (h *Handler) UseController(cc *controller.Controller, currentFc *config.FileConfig) {

	_c := func(v string) string {
		return checkPathWithSpace(h, v, "-")
	}

	for _, sheet := range *cc.Object.Sheets {

		var (
			url    string
			prefix string
			suffix string
		)

		if currentFc.Prefix != "" {
			prefix = utils.GenStandardPath(currentFc.Prefix)
			suffix = utils.GenStandardPath(sheet.Name)
			url = _c(utils.JoinStandardPath(prefix, suffix))
		} else {
			url = _c(utils.JoinStandardPath(cc.RawFilePath, sheet.Name))
		}

		// 保存数据
		h.HTMLInjectData = append(h.HTMLInjectData, HTMLInjectData{
			Url:  url,
			File: cc.RawFilePath,
		})

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
