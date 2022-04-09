package xlsx

import (
	"gitee.com/feimos/xs/common"
)

// Sheet
//
// 数据表结构体
type Sheet struct {
	Rows

	// 表的名称
	Name string

	// KEY 键名
	Key []string
}

func (s *Sheet) ToJson() common.JsonArray {
	var res common.JsonArray
	for _, row := range s.Rows {
		json := row.ToJson()
		res = append(res, *json)
	}
	return res
}
