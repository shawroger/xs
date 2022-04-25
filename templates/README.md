# templates

`import "gitee.com/feimos/xs/templates"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [func BindIndexHtml(params ...any) string](#func-bindindexhtml)
- [type StaticFile](#type-staticfile)
- [func ParseStaticFiles() []StaticFile](#func-parsestaticfiles)

## <a id="func-bindindexhtml">func</a> BindIndexHtml

```go
func BindIndexHtml(params ...any) string
```

BindIndexHtml

绑定 HTML 数据

## <a id="type-staticfile">type</a> StaticFile

```go
type StaticFile struct {
    Content string
    Url     string
}

```

### <a id="func-parsestaticfiles">func</a> ParseStaticFiles

```go
func ParseStaticFiles() []StaticFile
```

ParseStaticFiles

解析静态目录

---
