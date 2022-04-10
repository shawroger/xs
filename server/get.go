package server

import (
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/xlsx"
	"strconv"
)

// CreateServiceGet
//
// 创建 GET 请求服务
func (h *Handler) CreateServiceGet(url string, sheet *xlsx.Sheet) {

	h.debug("CREATE SERVICE HTTP::GET[%v]", url)

	h.GET(url, func(c *Context) {

		// 查找 ?_id=xxx 情况
		id := c.Query("_id")
		res := datatype.NewJson()

		if id != "" {
			json, err := handlerQueryID(id, sheet)

			switch err.(type) {
			case nil:
				{
					// 成功返回数据
					res.SetData(json)
					res.SetMsg("ok")
					break
				}
			case xlsx.ErrNotGetRowByIndex:
				{
					// 没有找到 index
					res.SetData([]string{})
					res.SetMsg("not found id %v", id)
					break
				}
			default:
				{
					// 不正确的 id 格式
					res.SetData([]string{})
					res.SetMsg("invalid id %v", id)
				}

			}

		} else {
			// 其他情况，返回所有
			res.SetData(sheet.ToJson())
			res.SetMsg("ok")
		}

		res.SetCode(200)
		c.JSON(200, res)
	})
}

// 处理有 _id 参数的情况
func handlerQueryID(id string, sheet *xlsx.Sheet) (*datatype.JsonMap, error) {

	index, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	// 输入的字符串，可以被解析为 int 类型
	row, err := sheet.GetRowByIndex(int(index))

	// 没有找到
	if err != nil {
		return nil, err
	}

	// 找到了对应的 Row
	json := row.ToJsonWithKeys(sheet.Key)
	json.SetID(int(index))
	return json, nil
}
