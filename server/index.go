package server

import (
	"gitee.com/feimos/xs/templates"
	"gitee.com/feimos/xs/utils"
	"github.com/gin-gonic/gin"
)

// CreateServiceIndex
//
// 创建首页请求服务
func (h *Handler) CreateServiceIndex() {

	h.GET("/", func(c *gin.Context) {
		html := templates.BindIndexHtml(h.HTMLInjectData, h.AppInfo.Version)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html)
	})
}

// CreateServiceStatic
//
// 创建 Static 请求服务
func (h *Handler) CreateServiceStatic() {
	files := templates.ParseStaticFiles()
	for _, file := range files {
		url := file.Url
		content := file.Content
		h.GET(url, func(c *gin.Context) {
			h.debug("Create static service at \"%v\"", url)
			c.Header("Content-Type", utils.GetMimeTypeFromFileName(url, true))
			c.String(200, content)
		})
	}

}
