package utils

// EqualAble
//
// 可以用 == 号的类型
type EqualAble interface {
	int | string | float32
}

// Exist
//
// 判断切片是否存在值相等元素
func Exist[T EqualAble](arr []T, v T) bool {
	for _, item := range arr {
		if item == v {
			return true
		}
	}

	return false
}
