package utils

import (
	"path"
	"regexp"
	"strings"
)

// UnifyPathFormat
//
// 文件路径名统一化：用 / 分隔
//
// 且路径被简化
func UnifyPathFormat(filename string) string {
	res := strings.Replace(filename, "\\", "/", -1)
	return path.Clean(res)
}

// GetPureName
//
// 获取纯单文件名 （无后缀名）
func GetPureName(filename string) string {
	filename = UnifyPathFormat(filename)
	extName := path.Ext(filename)
	baseName := path.Base(filename)
	return baseName[:len(baseName)-len(extName)]
}

// ToSnakeCase
//
// Camel 转 Snake
//
// 用 `-` 相连
func ToSnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

// GenStandardPath
//
// 生成标准路径字符串
func GenStandardPath(rawPath string) string {
	return UsePureFuncOneTypes(rawPath, GetPureName, ToSnakeCase)
}

// JoinStandardPath
//
// 连接多条路径，生成标准路径字符串
func JoinStandardPath(rawPaths ...string) string {
	root := GenStandardPath(rawPaths[0])
	for i, rawPath := range rawPaths {
		if i > 0 {
			root += "/" + GenStandardPath(rawPath)
		}
	}

	return root
}

// UnifyKeyName
//
// 键名 KEY 标准化
func UnifyKeyName(key string) string {
	key = GenStandardPath(key)
	key = strings.ReplaceAll(key, " ", "-")
	return key
}

// ParseCmdFilesFlag
//
// 解析 cmd.flags.files 字段
func ParseCmdFilesFlag(files string) []string {
	return strings.Split(files, "+")
}
