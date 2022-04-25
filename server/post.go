package server

import (
	"fmt"
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/xlsx"
	"strconv"
)

// CreateServicePost
//
// 创建 POST 请求服务
func (h *Handler) CreateServicePost(url string, f *xlsx.File, sheet *xlsx.Sheet) {

	h.debug("CREATE SERVICE HTTP::POST[%v]", url)

	h.POST(url, func(c *Context) {

		res := datatype.NewJson()

		// 查找 ?_id=xxx 情况
		id := c.Query("_id")

		// 获取所有参数列表
		valueList := getAllQueryKeys(c, sheet)

		// 设置了 _id，则为修改数据
		if id != "" {
			index, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				res.SetErrMsg(err.Error())
				res.SetMsg("edit fail")
			} else {
				err = setRowValue(f, int(index), sheet, valueList, false)
				if err != nil {
					res.SetErrMsg(err.Error())
					res.SetMsg("edit fail")
				} else {
					res.SetMsg("edit success")
					_ = sheet.EditRow(int(index), *xlsx.ParseRowFromValueList(valueList))
				}
			}

		} else {
			// 否则为新增数据
			err := setRowValue(f, 1+len(sheet.Rows), sheet, valueList, true)
			if err != nil {
				res.SetErrMsg(err.Error())
				res.SetMsg("add fail")
			} else {
				res.SetMsg("add success")
				sheet.AddRow(*xlsx.ParseRowFromValueList(valueList))
			}

		}

		res.SetCode(200)
		c.JSON(200, res)
	})
}

// 获取所有以 sheet.Key 为名的 query 值
func getAllQueryKeys(c *Context, sheet *xlsx.Sheet) []any {
	sheetLen := len(sheet.Key)
	res := make([]any, sheetLen)
	for index, key := range sheet.Key {
		query := c.Query(key)

		// 转为 Value
		val := xlsx.NewValue(query)
		res[index] = val.TrueValue

	}
	return res
}

// GenAxis
//
// 生成 index 所在行的 axis
func GenAxis[T int | int64 | uint](index T) string {
	return fmt.Sprintf("A%v", index+1)
}

// 设置行数据
func setRowValue(f *xlsx.File, index int, sheet *xlsx.Sheet, valueList []any, addFlag bool) error {

	// 不是新增数据
	if false == addFlag {
		row, err := sheet.GetRowByIndex(index)
		if err != nil {
			return err
		}
		// 混入数据
		MergeValueList(row, &valueList)
	}

	err := f.SetSheetRow(sheet.Name, GenAxis(index), &valueList)
	if err != nil {
		return err
	}

	return f.Save()
}

// MergeValueList
//
// 覆盖未定义部分
func MergeValueList(row *xlsx.Row, valueList *[]any) {
	for i, v := range *valueList {
		if (v == nil || v == "") && i < len(row.Val) {
			(*valueList)[i] = row.Val[i].TrueValue
		}
	}
}
