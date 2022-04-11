package utils

import (
	"strings"
)

// AutoNewLine
//
// 自动添加换行符
func AutoNewLine(text string) string {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}

	return text
}
