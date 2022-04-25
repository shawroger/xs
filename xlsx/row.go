package xlsx

import (
	"fmt"
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/utils"
)

// Row 单元格单行数据
//
// 从 Sheet 中读取到的若干行数据
type Row struct {

	// VAL 值
	Val []*Value
}

// Rows
//
// 多行数据
type Rows []Row

// ParseRowFromStringArr
//
// 从双重 string 数组创建 Row
func ParseRowFromStringArr(rows []string) *Row {
	var res Row
	for _, val := range rows {
		res.Val = append(res.Val, NewValue(val))
	}

	return &res
}

// ParseRowFromValueList
//
// 从双重 []any 数组创建 Row
func ParseRowFromValueList(valueList []any) *Row {
	var res Row
	for _, val := range valueList {
		res.Val = append(res.Val, NewValue(fmt.Sprintf("%v", val)))
	}

	return &res
}

// ToJsonWithKeys
//
// 将 Row 转为 json 格式
func (r *Row) ToJsonWithKeys(keys []string) *datatype.JsonMap {
	json := make(datatype.JsonMap)

	for i := 0; i < len(r.Val); i++ {
		if i < len(keys) {
			key := utils.UnifyKeyName(keys[i])
			val := r.Val[i].TrueValue
			json[key] = val
		}

	}
	return &json
}

// ToArray
//
// 将 Rows 转为 [][]string 格式
func (r *Rows) ToArray() [][]string {
	var res [][]string

	for _, row := range *r {
		item := make([]string, len(row.Val))
		for i, v := range row.Val {
			item[i] = v.Liter
		}
		res = append(res, item)
	}

	return res
}
