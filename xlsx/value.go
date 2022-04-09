package xlsx

import (
	"fmt"
	"strconv"
)

// ValueType
//
// 单元格元素数据可解析的值
type ValueType uint

// 打印 ValueType 代表的类型字面量
func (t ValueType) name() string {
	if t == 0 {
		return "int"
	} else if t == 1 {
		return "float"
	} else if t == 2 {
		return "string"
	} else {
		return "null"
	}
}

const (
	// ValueTypeInt
	// int 类型
	ValueTypeInt ValueType = iota
	// ValueTypeFloat
	// float 类型
	ValueTypeFloat ValueType = iota
	// ValueTypeString
	// string 类型
	ValueTypeString ValueType = iota
	// ValueTypeNull
	// 自定义 Null 类型
	ValueTypeNull ValueType = iota
)

// Value
// 单元格元素数据内容
type Value struct {
	// 字符串字面量
	Liter string

	// 实际可解析类型
	Type ValueType

	// 真实解析后的值
	TrueValue any
}

// NewValue
//
// 创建 Value 结构体
func NewValue(liter string) *Value {
	// 可以解析为 int
	if i, err := strconv.ParseInt(liter, 10, 64); err == nil {
		return &Value{
			TrueValue: i,
			Liter:     liter,
			Type:      ValueTypeInt,
		}
	}

	// 可以解析为 float
	if fl, err := strconv.ParseFloat(liter, 64); err == nil {
		return &Value{
			TrueValue: fl,
			Liter:     liter,
			Type:      ValueTypeFloat,
		}
	}

	return &Value{
		TrueValue: liter,
		Liter:     liter,
		Type:      ValueTypeString,
	}
}

// ErrValueConvert
//
// 自定义错误
//
// 转换 Value 类型错误
type ErrValueConvert struct {
	Val      any
	FromType ValueType
	IntoType ValueType
}

// 实现 error 接口
func (c ErrValueConvert) Error() string {
	return fmt.Sprintf("ErrValueConvert: Cannot convert value \"%v\" into type \"%v\", find type \"%v\"", c.Val, c.IntoType.name(), c.FromType.name())

}

// GenErrValueConvert 生成类型转换错误
func (v *Value) GenErrValueConvert(intoType ValueType) *ErrValueConvert {
	return &ErrValueConvert{
		Val:      v.TrueValue,
		FromType: v.Type,
		IntoType: intoType,
	}

}

func (v *Value) AsInt() (int, error) {
	if v.Type != ValueTypeInt {
		return 0, v.GenErrValueConvert(ValueTypeInt)
	}

	return v.TrueValue.(int), nil
}

func (v *Value) AsFloat() (float64, error) {
	if v.Type != ValueTypeFloat {
		return 0, v.GenErrValueConvert(ValueTypeFloat)
	}

	return v.TrueValue.(float64), nil
}

func (v *Value) AsString() (string, error) {
	if v.Type != ValueTypeString {
		return "", v.GenErrValueConvert(ValueTypeString)
	}

	return v.TrueValue.(string), nil
}
