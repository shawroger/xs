package datatype

// JsonMap
//
// 可用于输出的 json 类型
type JsonMap map[string]interface{}

// NewJson
//
// 创建一个空 JsonMap
func NewJson() *JsonMap {
	res := make(JsonMap, 0)
	return &res
}

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

// NewArray
//
// 创建一个空 JsonArray
func NewArray() *JsonArray {
	var res JsonArray = make([]JsonMap, 0)
	return &res
}

// Add
//
// 加入数据
func (a *JsonArray) Add(data JsonMap) {
	*a = append(*a, data)
}

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

// Unify
// JsonArray 转化为 JsonMap
//
// 数据在 data 字段内
func (a JsonArray) Unify() JsonMap {
	res := JsonMap{
		"data": a,
	}

	return res
}
