package utils

// UsePureFuncOneTypes
//
// 纯函数，同一类型单入单出
//
// 生成函数链
func UsePureFuncOneTypes[T any](v T, pureMapFunc ...func(v T) T) T {
	for _, fn := range pureMapFunc {
		v = fn(v)

	}
	return v
}
