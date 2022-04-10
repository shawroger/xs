package datatype

import "fmt"

// Set
//
// 设置键值
func (m *JsonMap) Set(key string, val any) {
	(*m)[key] = val
}

// SetID
//
// 设置内置 ID 值
func (m *JsonMap) SetID(val int) {
	m.Set("_id", val)
}

// SetData
//
// 设置内置 data 值
func (m *JsonMap) SetData(val any) {
	m.Set("data", val)
}

// SetCode
//
// 设置返回码
func (m *JsonMap) SetCode(code int) {
	m.Set("code", code)
}

// SetMsg
//
// 设置 msg 字符串
func (m *JsonMap) SetMsg(format string, values ...any) {
	var val string
	if len(values) == 0 {
		val = format
	} else {
		val = fmt.Sprintf(format, values)
	}

	m.Set("msg", val)
}

// Unify
