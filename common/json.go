package common

import "fmt"

// JsonMap
//
// 可用于输出的 json 类型
type JsonMap map[string]interface{}

// JsonUnify
//
// Json 类型标准化
type JsonUnify interface {
	Unify() JsonMap
}

// JsonArray
//
// json 中的数组
//
// 可用于输出的 json 类型
type JsonArray []JsonMap

// Unify
//
// 判断有无 data 字段
//
// 否则将所有数据放在 data 字段内
func (m JsonMap) Unify() JsonMap {
	if _, ok := m["data"]; ok {
		return m
	} else {
		return JsonMap{
			"data": m,
		}
	}
}

// SetCode
//
// 设置返回码
func (m *JsonMap) SetCode(code int) {
	(*m)["code"] = code
}

// SetMsg
//
// 设置 msg 字符串
func (m *JsonMap) SetMsg(format string, values ...any) {
	if len(values) == 0 {
		(*m)["msg"] = format
	} else {
		(*m)["msg"] = fmt.Sprintf(format, values)
	}
}

// Unify
//
// JsonArray 转化为 JsonMap
//
// 数据在 data 字段内
func (a JsonArray) Unify() JsonMap {
	res := JsonMap{
		"data": a,
	}

	return res
}
