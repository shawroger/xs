package server

import (
	"gitee.com/feimos/xs/config"
	"gitee.com/feimos/xs/datatype"
	"gitee.com/feimos/xs/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// HTMLInjectData
//
// 注入 index.html 所需的数据
type HTMLInjectData struct {
	Url  string
	File string
}

// Handler
//
// 包含 *gin.Engine 所有方法
//
// 服务器主引擎
type Handler struct {
	*gin.Engine
	Debug          bool
	GinMode        bool
	Port           string
	HTMLInjectData []HTMLInjectData
	Config         *config.Config
	AppInfo        *datatype.AppInfo
}

// New
//
// 创建服务器 Handler
//
// 即创建 gin 服务器
func New() *Handler {
	h := &Handler{}
	h.Port = ":8080"
	h.EnableDebugMode()
	h.DisableGinDebugMode()
	h.Engine = gin.Default()

	// 开启静态资源服务
	h.CreateServiceStatic()

	return h
}

// EnableDebugMode
//
// Debug 模式 开启通知
func (h *Handler) EnableDebugMode() {

	h.Debug = true

}

// DisableDebugMode
//
// 不开启 Debug 减少日志通知
func (h *Handler) DisableDebugMode() {

	h.Debug = false
}

// EnableGinDebugMode
//
// 开启 gin 框架的 Debug 模式
func (h *Handler) EnableGinDebugMode() {
	// 关闭 gin 通知
	gin.SetMode(gin.DebugMode)
}

// DisableGinDebugMode
//
// 不开启 gin 框架的 Debug 模式
func (h *Handler) DisableGinDebugMode() {
	// 关闭 gin 通知
	gin.SetMode(gin.ReleaseMode)
}

// Run
//
// 启动服务器
func (h *Handler) Run() error {
	c, err := h.GetConfig()
	if err != nil {
		return err
	}

	// 覆盖 Port
	if c.Port != 0 {
		h.Port = ":" + strconv.Itoa(c.Port)
	}

	h.log("Server is running at http://localhost%v", h.Port)
	return h.Engine.Run(h.Port)
}

// 打印 debug 日志
func (h *Handler) debug(format string, values ...any) {
	if h.Debug {
		log.SetPrefix("[XS-DEBUG]")
		log.Printf(utils.AutoNewLine(format), values...)
	}
}

// 打印普通日志
func (h *Handler) log(format string, values ...any) {

	log.SetPrefix("[XS]")
	log.Printf(utils.AutoNewLine(format), values...)

}
