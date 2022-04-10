package controller

import (
	"gitee.com/feimos/xs/datatype"
)

// Reader
//
// 读取数据应当实现的接口
type Reader interface {
	// ToJson
	// 转为 json 格式
	ToJson() *datatype.JsonMap
}
