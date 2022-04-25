# controller

`import "gitee.com/feimos/xs/controller"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [type Controller](#type-controller)
- [func OpenFile(filename string) (\*Controller, error)](#func-openfile)
  - [func (c Controller) GenPrefixURL() string](#func-controller-genprefixurl)
- [type Reader](#type-reader)

## <a id="type-controller">type</a> Controller

```go
type Controller struct {
    // 文件操作器
    *excelize.File

    // 数据对象
    *xlsx.Object

    // 原始文件路径
    RawFilePath string
}

```

Controller

控制器结构体

### <a id="func-openfile">func</a> OpenFile

```go
func OpenFile(filename string) (*Controller, error)
```

OpenFile

读取文件，从文件中创建 Controller

### <a id="func-controller-genprefixurl">func</a> (Controller) GenPrefixURL

```go
func (c Controller) GenPrefixURL() string
```

GenPrefixURL

生成文件前缀 URL 字符串

## <a id="type-reader">type</a> Reader

```go
type Reader interface {
    // ToJson
    // 转为 json 格式
    ToJson() *datatype.JsonMap
}
```

Reader

读取数据应当实现的接口

---
