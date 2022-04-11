package xlsx

import (
	"fmt"
	"gitee.com/feimos/xs/datatype"
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

// ToJson
//
// 将 Sheet 转为 json Array 格式
func (s *Sheet) ToJson() datatype.JsonArray {
	var res datatype.JsonArray
	for i, row := range s.Rows {
		json := row.ToJsonWithKeys(s.Key)
		json.SetID(i + 1)
		res = append(res, *json)
	}
	return res
}

// ErrNotGetIndexByKey
//
// 自定义错误类型
//
// GetIndexByKey 返回的错误
type ErrNotGetIndexByKey struct {
	key  string
	name string
}

func (e ErrNotGetIndexByKey) Error() string {
	return fmt.Sprintf("Can not find key %v in sheet %v", e.key, e.name)
}

// GetIndexByKey
//
// 根据 KEY 名称获取 axis 序号
func (s *Sheet) GetIndexByKey(key string) (int, error) {

	for i, k := range s.Key {
		if k == key {
			return i, nil
		}
	}

	return 0, ErrNotGetIndexByKey{key: key, name: s.Name}
}

// ErrNotGetRowByIndex
//
// 自定义错误类型
//
// GetRowByIndex 返回的错误
type ErrNotGetRowByIndex struct {
	Index int
	Name  string
}

func (e ErrNotGetRowByIndex) Error() string {
	return fmt.Sprintf("Can not find index %v in sheet %v", e.Index, e.Name)
}

// GetRowByIndex
//
// 通过 index 获取 row
func (s *Sheet) GetRowByIndex(index int) (*Row, error) {
	for i, row := range s.Rows {
		if i == index-1 {
			return &row, nil
		}
	}

	return nil, ErrNotGetRowByIndex{
		Index: index,
		Name:  s.Name,
	}
}
