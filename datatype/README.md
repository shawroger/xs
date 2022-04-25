# datatype

`import "gitee.com/feimos/xs/datatype"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [type AppInfo](#type-appinfo)
- [type JsonArray](#type-jsonarray)
- [func NewArray() \*JsonArray](#func-newarray)
  - [func (a \*JsonArray) Add(data JsonMap)](#func-jsonarray-add)
  - [func (a JsonArray) Unify() JsonMap](#func-jsonarray-unify)
- [type JsonMap](#type-jsonmap)
- [func NewJson() \*JsonMap](#func-newjson)
  - [func (m \*JsonMap) Set(key string, val any)](#func-jsonmap-set)
  - [func (m \*JsonMap) SetCode(code int)](#func-jsonmap-setcode)
  - [func (m \*JsonMap) SetData(val any)](#func-jsonmap-setdata)
  - [func (m \*JsonMap) SetErrMsg(format string, values ...any)](#func-jsonmap-seterrmsg)
  - [func (m \*JsonMap) SetID(val int)](#func-jsonmap-setid)
  - [func (m \*JsonMap) SetMsg(format string, values ...any)](#func-jsonmap-setmsg)
  - [func (m JsonMap) Unify() JsonMap](#func-jsonmap-unify)
- [type JsonUnify](#type-jsonunify)

## <a id="type-appinfo">type</a> AppInfo

```go
type AppInfo struct {
    Version string
}

```

AppInfo

配置 cmd 信息

## <a id="type-jsonarray">type</a> JsonArray

```go
type JsonArray []JsonMap
```

JsonArray

json 中的数组

可用于输出的 json 类型

### <a id="func-newarray">func</a> NewArray

```go
func NewArray() *JsonArray
```

NewArray

创建一个空 JsonArray

### <a id="func-jsonarray-add">func</a> (\*JsonArray) Add

```go
func (a *JsonArray) Add(data JsonMap)
```

Add

加入数据

### <a id="func-jsonarray-unify">func</a> (JsonArray) Unify

```go
func (a JsonArray) Unify() JsonMap
```

Unify
JsonArray 转化为 JsonMap

数据在 data 字段内

## <a id="type-jsonmap">type</a> JsonMap

```go
type JsonMap map[string]interface{}
```

JsonMap

可用于输出的 json 类型

### <a id="func-newjson">func</a> NewJson

```go
func NewJson() *JsonMap
```

NewJson

创建一个空 JsonMap

### <a id="func-jsonmap-set">func</a> (\*JsonMap) Set

```go
func (m *JsonMap) Set(key string, val any)
```

Set

设置键值

### <a id="func-jsonmap-setcode">func</a> (\*JsonMap) SetCode

```go
func (m *JsonMap) SetCode(code int)
```

SetCode

设置返回码

### <a id="func-jsonmap-setdata">func</a> (\*JsonMap) SetData

```go
func (m *JsonMap) SetData(val any)
```

SetData

设置内置 data 值

### <a id="func-jsonmap-seterrmsg">func</a> (\*JsonMap) SetErrMsg

```go
func (m *JsonMap) SetErrMsg(format string, values ...any)
```

SetErrMsg

设置 error 字符串

### <a id="func-jsonmap-setid">func</a> (\*JsonMap) SetID

```go
func (m *JsonMap) SetID(val int)
```

SetID

设置内置 ID 值

### <a id="func-jsonmap-setmsg">func</a> (\*JsonMap) SetMsg

```go
func (m *JsonMap) SetMsg(format string, values ...any)
```

SetMsg

设置 msg 字符串

### <a id="func-jsonmap-unify">func</a> (JsonMap) Unify

```go
func (m JsonMap) Unify() JsonMap
```

Unify

判断有无 data 字段

否则将所有数据放在 data 字段内

## <a id="type-jsonunify">type</a> JsonUnify

```go
type JsonUnify interface {
    Unify() JsonMap
}
```

JsonUnify

Json 类型标准化

---
