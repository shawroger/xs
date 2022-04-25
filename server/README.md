# server

`import "gitee.com/feimos/xs/server"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [func GenAxis[T int | int64 | uint](index T) string](#func-genaxis)
- [func HandlerCreateService(h *Handler, url string, f *xlsx.File, sheet xlsx.Sheet)](#func-handlercreateservice)
- [func MergeValueList(row _xlsx.Row, valueList _[]any)](#func-mergevaluelist)
- [type Context](#type-context)
- [type ErrMissingHandlerConfig](#type-errmissinghandlerconfig)
  - [func (c ErrMissingHandlerConfig) Error() string](#func-errmissinghandlerconfig-error)
- [type HTMLInjectData](#type-htmlinjectdata)
- [type Handler](#type-handler)
- [func New() \*Handler](#func-new)
  - [func (h *Handler) BindAppInfo(a *datatype.AppInfo) \*Handler](#func-handler-bindappinfo)
  - [func (h *Handler) CreateServiceDelete(url string, f *xlsx.File, sheet \*xlsx.Sheet)](#func-handler-createservicedelete)
  - [func (h *Handler) CreateServiceGet(url string, sheet *xlsx.Sheet)](#func-handler-createserviceget)
  - [func (h \*Handler) CreateServiceIndex()](#func-handler-createserviceindex)
  - [func (h *Handler) CreateServicePost(url string, f *xlsx.File, sheet \*xlsx.Sheet)](#func-handler-createservicepost)
  - [func (h \*Handler) CreateServiceStatic()](#func-handler-createservicestatic)
  - [func (h \*Handler) DisableDebugMode()](#func-handler-disabledebugmode)
  - [func (h \*Handler) DisableGinDebugMode()](#func-handler-disablegindebugmode)
  - [func (h \*Handler) EnableDebugMode()](#func-handler-enabledebugmode)
  - [func (h \*Handler) EnableGinDebugMode()](#func-handler-enablegindebugmode)
  - [func (h *Handler) GetConfig() (*config.Config, error)](#func-handler-getconfig)
  - [func (h \*Handler) LoadConfigFile(filename string) error](#func-handler-loadconfigfile)
  - [func (h \*Handler) Run() error](#func-handler-run)
  - [func (h *Handler) UseConfig(c *config.Config) error](#func-handler-useconfig)
  - [func (h *Handler) UseController(cc *controller.Controller, fileConfig \*config.FileConfig)](#func-handler-usecontroller)

## <a id="func-genaxis">func</a> GenAxis

```go
func GenAxis[T int | int64 | uint](index T) string
```

GenAxis

生成 index 所在行的 axis

## <a id="func-handlercreateservice">func</a> HandlerCreateService

```go
func HandlerCreateService(h *Handler, url string, f *xlsx.File, sheet xlsx.Sheet)
```

HandlerCreateService

创建服务

必须传入的是 sheet 而不能是 &sheet

## <a id="func-mergevaluelist">func</a> MergeValueList

```go
func MergeValueList(row *xlsx.Row, valueList *[]any)
```

MergeValueList

覆盖未定义部分

## <a id="type-context">type</a> Context

```go
type Context = gin.Context
```

Context

gin.Context 的同名

## <a id="type-errmissinghandlerconfig">type</a> ErrMissingHandlerConfig

```go
type ErrMissingHandlerConfig struct {
}

```

ErrMissingHandlerConfig

自定义错误：未定义配置

### <a id="func-errmissinghandlerconfig-error">func</a> (ErrMissingHandlerConfig) Error

```go
func (c ErrMissingHandlerConfig) Error() string
```

## <a id="type-htmlinjectdata">type</a> HTMLInjectData

```go
type HTMLInjectData struct {
    Url  string
    File string
}

```

HTMLInjectData

注入 index.html 所需的数据

## <a id="type-handler">type</a> Handler

```go
type Handler struct {
    *gin.Engine
    Debug          bool
    GinMode        bool
    Port           string
    HTMLInjectData []HTMLInjectData
    Config         *config.Config
    AppInfo        *datatype.AppInfo
}

```

Handler

包含 \*gin.Engine 所有方法

服务器主引擎

### <a id="func-new">func</a> New

```go
func New() *Handler
```

New

创建服务器 Handler

即创建 gin 服务器

### <a id="func-handler-bindappinfo">func</a> (\*Handler) BindAppInfo

```go
func (h *Handler) BindAppInfo(a *datatype.AppInfo) *Handler
```

BindAppInfo

从 cmd 中绑定 AppInfo

### <a id="func-handler-createservicedelete">func</a> (\*Handler) CreateServiceDelete

```go
func (h *Handler) CreateServiceDelete(url string, f *xlsx.File, sheet *xlsx.Sheet)
```

CreateServiceDelete

创建 DELETE 请求服务

### <a id="func-handler-createserviceget">func</a> (\*Handler) CreateServiceGet

```go
func (h *Handler) CreateServiceGet(url string, sheet *xlsx.Sheet)
```

CreateServiceGet

创建 GET 请求服务

### <a id="func-handler-createserviceindex">func</a> (\*Handler) CreateServiceIndex

```go
func (h *Handler) CreateServiceIndex()
```

CreateServiceIndex

创建首页请求服务

### <a id="func-handler-createservicepost">func</a> (\*Handler) CreateServicePost

```go
func (h *Handler) CreateServicePost(url string, f *xlsx.File, sheet *xlsx.Sheet)
```

CreateServicePost

创建 POST 请求服务

### <a id="func-handler-createservicestatic">func</a> (\*Handler) CreateServiceStatic

```go
func (h *Handler) CreateServiceStatic()
```

CreateServiceStatic

创建 Static 请求服务

### <a id="func-handler-disabledebugmode">func</a> (\*Handler) DisableDebugMode

```go
func (h *Handler) DisableDebugMode()
```

DisableDebugMode

不开启 Debug 减少日志通知

### <a id="func-handler-disablegindebugmode">func</a> (\*Handler) DisableGinDebugMode

```go
func (h *Handler) DisableGinDebugMode()
```

DisableGinDebugMode

不开启 gin 框架的 Debug 模式

### <a id="func-handler-enabledebugmode">func</a> (\*Handler) EnableDebugMode

```go
func (h *Handler) EnableDebugMode()
```

EnableDebugMode

Debug 模式 开启通知

### <a id="func-handler-enablegindebugmode">func</a> (\*Handler) EnableGinDebugMode

```go
func (h *Handler) EnableGinDebugMode()
```

EnableGinDebugMode

开启 gin 框架的 Debug 模式

### <a id="func-handler-getconfig">func</a> (\*Handler) GetConfig

```go
func (h *Handler) GetConfig() (*config.Config, error)
```

GetConfig

读取配置

### <a id="func-handler-loadconfigfile">func</a> (\*Handler) LoadConfigFile

```go
func (h *Handler) LoadConfigFile(filename string) error
```

LoadConfigFile

从配置文件中加载配置

### <a id="func-handler-run">func</a> (\*Handler) Run

```go
func (h *Handler) Run() error
```

Run

启动服务器

### <a id="func-handler-useconfig">func</a> (\*Handler) UseConfig

```go
func (h *Handler) UseConfig(c *config.Config) error
```

UseConfig

加载配置

### <a id="func-handler-usecontroller">func</a> (\*Handler) UseController

```go
func (h *Handler) UseController(cc *controller.Controller, fileConfig *config.FileConfig)
```

UseController

通过加载 \*controller.Controller 创建服务器

currentFc 为当前加载的文件配置

---
