# utils

`import "gitee.com/feimos/xs/utils"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [func AutoNewLine(text string) string](#func-autonewline)
- [func Exist[T EqualAble](arr []T, v T) bool](#func-exist)
- [func GenStandardPath(rawPath string) string](#func-genstandardpath)
- [func GetMimeTypeFromExt(ext string) string](#func-getmimetypefromext)
- [func GetMimeTypeFromFileName(name string, charsetUTF8 bool) string](#func-getmimetypefromfilename)
- [func GetPureName(filename string) string](#func-getpurename)
- [func JoinStandardPath(rawPaths ...string) string](#func-joinstandardpath)
- [func ParseCmdFilesFlag(files string) ([]string, error)](#func-parsecmdfilesflag)
- [func ToSnakeCase(s string) string](#func-tosnakecase)
- [func UnifyKeyName(key string) string](#func-unifykeyname)
- [func UnifyPathFormat(filename string) string](#func-unifypathformat)
- [func UsePureFuncOneTypes[T any](v T, pureMapFunc ...func(v T) T) T](#func-usepurefunconetypes)
- [type EqualAble](#type-equalable)

## <a id="func-autonewline">func</a> AutoNewLine

```go
func AutoNewLine(text string) string
```

AutoNewLine

自动添加换行符

## <a id="func-exist">func</a> Exist

```go
func Exist[T EqualAble](arr []T, v T) bool
```

Exist

判断切片是否存在值相等元素

## <a id="func-genstandardpath">func</a> GenStandardPath

```go
func GenStandardPath(rawPath string) string
```

GenStandardPath

生成标准路径字符串

## <a id="func-getmimetypefromext">func</a> GetMimeTypeFromExt

```go
func GetMimeTypeFromExt(ext string) string
```

GetMimeTypeFromExt

从文件后缀名获取正确的 MIMETYPE

## <a id="func-getmimetypefromfilename">func</a> GetMimeTypeFromFileName

```go
func GetMimeTypeFromFileName(name string, charsetUTF8 bool) string
```

GetMimeTypeFromFileName

从文件名获取正确的 MIMETYPE

设置 charset 为 true 则增加 `charset=utf-8`

## <a id="func-getpurename">func</a> GetPureName

```go
func GetPureName(filename string) string
```

GetPureName

获取纯单文件名 （无后缀名）

## <a id="func-joinstandardpath">func</a> JoinStandardPath

```go
func JoinStandardPath(rawPaths ...string) string
```

JoinStandardPath

连接多条路径，生成标准路径字符串

## <a id="func-parsecmdfilesflag">func</a> ParseCmdFilesFlag

```go
func ParseCmdFilesFlag(files string) ([]string, error)
```

ParseCmdFilesFlag

解析 cmd.flags.files 字段

## <a id="func-tosnakecase">func</a> ToSnakeCase

```go
func ToSnakeCase(s string) string
```

ToSnakeCase

### Camel 转 Snake

用 `-` 相连

## <a id="func-unifykeyname">func</a> UnifyKeyName

```go
func UnifyKeyName(key string) string
```

UnifyKeyName

键名 KEY 标准化

## <a id="func-unifypathformat">func</a> UnifyPathFormat

```go
func UnifyPathFormat(filename string) string
```

UnifyPathFormat

文件路径名统一化：用 / 分隔

且路径被简化

## <a id="func-usepurefunconetypes">func</a> UsePureFuncOneTypes

```go
func UsePureFuncOneTypes[T any](v T, pureMapFunc ...func(v T) T) T
```

UsePureFuncOneTypes

纯函数，同一类型单入单出

生成函数链

## <a id="type-equalable">type</a> EqualAble

```go
type EqualAble interface {
    int | string | float32
}
```

EqualAble

可以用 == 号的类型

---
