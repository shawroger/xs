package server

import (
	"fmt"
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

		res := datatype.NewJson()

		// 查找 ?_id=xxx 情况
		id := c.Query("_id")

		// 存在 query[_id]
		if id != "" {
			json, err := parseID(id, sheet)
			handlerErrType(res, json, err, id)

		} else {

			// 获取所有参数列表
			valueList := getAllQueryKeys(c, sheet)
			if isMissingValueList(valueList) {
				json := sheet.ToJson()
				res.SetData(json)
				res.SetMsg("ok")
			} else {
				jsonArray := parseValueList(valueList, sheet)
				res.SetData(jsonArray)
				res.SetMsg("ok")
			}

		}

		res.SetCode(200)
		c.JSON(200, res)
	})
}

// 处理有其他参数的情况
func parseValueList(valueList []any, sheet *xlsx.Sheet) *datatype.JsonArray {

	jsonList := datatype.NewArray()
	for index, row := range sheet.Rows {
		rowValue := row.Val
		for i, val := range rowValue {

			if valueList[i] == val.TrueValue && val.Liter != "" {
				// 找到了对应的 Row
				json := row.ToJsonWithKeys(sheet.Key)
				json.SetID(int(index))
				jsonList.Add(*json)
			}

		}
	}
	return jsonList
}

// 判断是不是没有传参
func isMissingValueList(valueList []any) bool {
	for _, v := range valueList {
		if fmt.Sprintf("%v", v) != "" {
			return false
		}
	}

	return true
}

// 处理有 _id 参数的情况
func parseID(id string, sheet *xlsx.Sheet) (*datatype.JsonMap, error) {

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

// 处理不同 err 类型情况下的 json 值
func handlerErrType(res, json *datatype.JsonMap, err error, id string) {

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
			res.SetData(datatype.NewArray())
			res.SetMsg("not found id %v", id)
			break
		}
	default:
		{
			// 不正确的 id 格式
			res.SetData(datatype.NewArray())
			res.SetMsg("invalid id %v", id)
		}

	}
}
