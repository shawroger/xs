package xlsx

import (
	"gitee.com/feimos/xs/common"
	"gitee.com/feimos/xs/utils"
)

// Row 单元格单行数据
//
// 从 Sheet 中读取到的若干行数据
type Row struct {
	// 单行长度
	Len int

	// KEY 键名
	Key []string

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
func ParseRowFromStringArr(keys, rows []string) *Row {
	var res Row
	for _, val := range rows {
		res.Val = append(res.Val, NewValue(val))

	}

	// 设置键名
	res.Key = keys

	// 设置长度
	// 取较小值
	keyLen := len(res.Key)
	valLen := len(res.Val)

	if keyLen < valLen {
		res.Len = keyLen
	} else {
		res.Len = valLen
	}

	return &res
}

// ToJson
//
// 将 Row 转为 json 格式
func (r *Row) ToJson() *common.JsonMap {
	json := make(common.JsonMap)

	for i := 0; i < r.Len; i++ {
		key := utils.UnifyKeyName(r.Key[i])
		val := r.Val[i].TrueValue
		json[key] = val
	}

	return &json
}
