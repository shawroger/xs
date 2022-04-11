package server

import (
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/xlsx"
	"strconv"
)

// CreateServiceDelete
//
// 创建 DELETE 请求服务
func (h *Handler) CreateServiceDelete(url string, f *xlsx.File, sheet *xlsx.Sheet) {

	h.debug("CREATE SERVICE HTTP::DELETE[%v]", url)

	h.DELETE(url, func(c *Context) {

		res := datatype.NewJson()

		// 查找 ?_id=xxx 情况
		id := c.Query("_id")

		// 设置了 _id
		if id != "" {
			index, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				res.SetErrMsg(err.Error())
				res.SetMsg("delete fail")
			} else {
				err = deleteRow(f, int(index), sheet)
				if err != nil {
					res.SetErrMsg(err.Error())
					res.SetMsg("delete fail")
				} else {
					res.SetMsg("delete success")
				}
			}

		} else {
			res.SetErrMsg("missing id")
			res.SetMsg("delete fail")

		}

		res.SetCode(200)
		c.JSON(200, res)
	})
}

// 删除行数据
func deleteRow(f *xlsx.File, index int, sheet *xlsx.Sheet) error {
	_, err := sheet.GetRowByIndex(index)
	if err != nil {
		return err
	}

	err = f.RemoveRow(sheet.Name, index+1)
	if err != nil {
		return err
	}

	return f.Save()
}
