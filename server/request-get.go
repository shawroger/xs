package server

import (
	"gitee.com/feimos/xs/xlsx"
	"github.com/gin-gonic/gin"
)

// SetRequestGet
//
// 创建 GET 请求服务
func (h *Handler) SetRequestGet(url string, sheet *xlsx.Sheet) {

	h.debug("CREATE HTTP::GET[%v]\n", url)

	h.GET(url, func(c *gin.Context) {

		json := sheet.ToJson().Unify()
		json.SetCode(200)
		json.SetMsg("ok")

		c.JSON(200, json)
	})
}
