package controller

import (
	"gitee.com/feimos/xs/utils"
	"gitee.com/feimos/xs/xlsx"
	"github.com/xuri/excelize/v2"
)

// Controller
//
// 控制器结构体
type Controller struct {
	// 文件操作器
	*excelize.File

	// 数据对象
	*xlsx.Object

	// 原始文件路径
	RawFilePath string
}

// GenPrefixURL
//
// 生成文件前缀 URL 字符串
func (c Controller) GenPrefixURL() string {
	prefix := utils.GetPureName(c.RawFilePath)
	return utils.ToSnakeCase(prefix)
}

// OpenFile
//
// 读取文件，从文件中创建 Controller
func OpenFile(filename string) (*Controller, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	// 关闭文件
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	object, err := xlsx.NewObjectFromFile(file)
	if err != nil {
		return nil, err
	}

	// 设置 FileName

	return &Controller{
		File:        file,
		Object:      object,
		RawFilePath: filename,
	}, nil
}
