package utils

import (
	"path"
	"strings"
)

// GetMimeTypeFromExt
//
// 从文件后缀名获取正确的 MIMETYPE
func GetMimeTypeFromExt(ext string) string {
	if ext == ".css" {
		return "text/css"
	}

	if ext == ".js" {
		return "text/javascript"
	}

	if ext == ".html" {
		return "text/html"
	}

	return "text/plain"
}

// GetMimeTypeFromFileName
//
// 从文件名获取正确的 MIMETYPE
//
// 设置 charset 为 true 则增加 `charset=utf-8`
func GetMimeTypeFromFileName(name string, charsetUTF8 bool) string {
	var suffix string
	if charsetUTF8 {
		suffix = "; charset=utf-8"
	}
	ext := strings.ToLower(path.Ext(name))
	return GetMimeTypeFromExt(ext) + suffix
}
