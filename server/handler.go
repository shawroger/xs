package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// Handler
//
// 包含 *gin.Engine 所有方法
//
// 服务器主引擎
type Handler struct {
	*gin.Engine
	Debug   bool
	GinMode bool
}

// New
//
// 创建服务器 Handler
//
// 即创建 gin 服务器
func New() *Handler {
	// 关闭 gin 通知
	gin.SetMode(gin.ReleaseMode)

	h := &Handler{}
	h.Debug = false
	h.Engine = gin.Default()
	return h
}

// 打印 debug 日志
func (h *Handler) debug(format string, values ...any) {
	if h.Debug {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		log.SetPrefix("[XS-DEBUG]")
		log.Printf(format, values...)
	}
}

// 打印普通日志
func (h *Handler) log(format string, values ...any) {

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	log.SetPrefix("[XS-LOG]")
	log.Printf(format, values...)

}
