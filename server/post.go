package server

import "gitee.com/feimos/xs/xlsx"

// CreateServicePost
//
// 创建 POST 请求服务
func (h *Handler) CreateServicePost(url string, file *xlsx.File, sheet *xlsx.Sheet) {

	h.debug("CREATE SERVICE HTTP::POST[%v]", url)

	h.POST(url, func(c *Context) {

		json := sheet.ToJson().Unify()
		json.SetCode(200)
		json.SetMsg("ok")

		c.JSON(200, json)
	})
}
