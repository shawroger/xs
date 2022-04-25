# xlsx

`import "gitee.com/feimos/xs/xlsx"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [type ErrNotGetIndexByKey](#type-errnotgetindexbykey)
  - [func (e ErrNotGetIndexByKey) Error() string](#func-errnotgetindexbykey-error)
- [type ErrNotGetRowByIndex](#type-errnotgetrowbyindex)
  - [func (e ErrNotGetRowByIndex) Error() string](#func-errnotgetrowbyindex-error)
- [type ErrValueConvert](#type-errvalueconvert)
  - [func (c ErrValueConvert) Error() string](#func-errvalueconvert-error)
- [type File](#type-file)
- [type Object](#type-object)
- [func NewObjectFromFile(f *File) (*Object, error)](#func-newobjectfromfile)
- [type Row](#type-row)
- [func ParseRowFromStringArr(rows []string) \*Row](#func-parserowfromstringarr)
- [func ParseRowFromValueList(valueList []any) \*Row](#func-parserowfromvaluelist)
  - [func (r *Row) ToJsonWithKeys(keys []string) *datatype.JsonMap](#func-row-tojsonwithkeys)
- [type Rows](#type-rows)
  - [func (r \*Rows) ToArray() [][]string](#func-rows-toarray)
- [type Sheet](#type-sheet)
  - [func (s \*Sheet) AddRow(row Row)](#func-sheet-addrow)
  - [func (s \*Sheet) EditRow(i int, row Row) error](#func-sheet-editrow)
  - [func (s \*Sheet) GetIndexByKey(key string) (int, error)](#func-sheet-getindexbykey)
  - [func (s *Sheet) GetRowByIndex(index int) (*Row, error)](#func-sheet-getrowbyindex)
  - [func (s \*Sheet) RemoveRowByIndex(i int) error](#func-sheet-removerowbyindex)
  - [func (s \*Sheet) ToJson() datatype.JsonArray](#func-sheet-tojson)
- [type Value](#type-value)
- [func NewValue(liter string) \*Value](#func-newvalue)
  - [func (v \*Value) AsFloat() (float64, error)](#func-value-asfloat)
  - [func (v \*Value) AsInt() (int, error)](#func-value-asint)
  - [func (v \*Value) AsString() (string, error)](#func-value-asstring)
  - [func (v *Value) GenErrValueConvert(intoType ValueType) *ErrValueConvert](#func-value-generrvalueconvert)
- [type ValueType](#type-valuetype)

## <a id="type-errnotgetindexbykey">type</a> ErrNotGetIndexByKey

```go
type ErrNotGetIndexByKey struct {
    // contains filtered or unexported fields
}

```

ErrNotGetIndexByKey

自定义错误类型

GetIndexByKey 返回的错误

### <a id="func-errnotgetindexbykey-error">func</a> (ErrNotGetIndexByKey) Error

```go
func (e ErrNotGetIndexByKey) Error() string
```

## <a id="type-errnotgetrowbyindex">type</a> ErrNotGetRowByIndex

```go
type ErrNotGetRowByIndex struct {
    Index int
    Name  string
}

```

ErrNotGetRowByIndex

自定义错误类型

GetRowByIndex 返回的错误

### <a id="func-errnotgetrowbyindex-error">func</a> (ErrNotGetRowByIndex) Error

```go
func (e ErrNotGetRowByIndex) Error() string
```

## <a id="type-errvalueconvert">type</a> ErrValueConvert

```go
type ErrValueConvert struct {
    Val      any
    FromType ValueType
    IntoType ValueType
}

```

ErrValueConvert

自定义错误

转换 Value 类型错误

### <a id="func-errvalueconvert-error">func</a> (ErrValueConvert) Error

```go
func (c ErrValueConvert) Error() string
```

实现 error 接口

## <a id="type-file">type</a> File

```go
type File = excelize.File
```

File

excelize.File 别名

## <a id="type-object">type</a> Object

```go
type Object struct {
    Sheets *[]Sheet
}

```

Object

对于每一个 XLSX 文件

生成一个独立的 Object 对象

### <a id="func-newobjectfromfile">func</a> NewObjectFromFile

```go
func NewObjectFromFile(f *File) (*Object, error)
```

NewObjectFromFile

从 \*excelize.File 句柄生成 Object 对象

## <a id="type-row">type</a> Row

```go
type Row struct {

    // VAL 值
    Val []*Value
}

```

Row 单元格单行数据

从 Sheet 中读取到的若干行数据

### <a id="func-parserowfromstringarr">func</a> ParseRowFromStringArr

```go
func ParseRowFromStringArr(rows []string) *Row
```

ParseRowFromStringArr

从双重 string 数组创建 Row

### <a id="func-parserowfromvaluelist">func</a> ParseRowFromValueList

```go
func ParseRowFromValueList(valueList []any) *Row
```

ParseRowFromValueList

从双重 []any 数组创建 Row

### <a id="func-row-tojsonwithkeys">func</a> (\*Row) ToJsonWithKeys

```go
func (r *Row) ToJsonWithKeys(keys []string) *datatype.JsonMap
```

ToJsonWithKeys

将 Row 转为 json 格式

## <a id="type-rows">type</a> Rows

```go
type Rows []Row
```

Rows

多行数据

### <a id="func-rows-toarray">func</a> (\*Rows) ToArray

```go
func (r *Rows) ToArray() [][]string
```

ToArray

将 Rows 转为 [][]string 格式

## <a id="type-sheet">type</a> Sheet

```go
type Sheet struct {
    Rows

    // 表的名称
    Name string

    // KEY 键名
    Key []string
}

```

Sheet

数据表结构体

### <a id="func-sheet-addrow">func</a> (\*Sheet) AddRow

```go
func (s *Sheet) AddRow(row Row)
```

AddRow

在末尾新增一项 row

### <a id="func-sheet-editrow">func</a> (\*Sheet) EditRow

```go
func (s *Sheet) EditRow(i int, row Row) error
```

EditRow

在指定 index 处修改 row

### <a id="func-sheet-getindexbykey">func</a> (\*Sheet) GetIndexByKey

```go
func (s *Sheet) GetIndexByKey(key string) (int, error)
```

GetIndexByKey

根据 KEY 名称获取 axis 序号

### <a id="func-sheet-getrowbyindex">func</a> (\*Sheet) GetRowByIndex

```go
func (s *Sheet) GetRowByIndex(index int) (*Row, error)
```

GetRowByIndex

通过 index 获取 row

### <a id="func-sheet-removerowbyindex">func</a> (\*Sheet) RemoveRowByIndex

```go
func (s *Sheet) RemoveRowByIndex(i int) error
```

RemoveRowByIndex

通过 index 删除 row

### <a id="func-sheet-tojson">func</a> (\*Sheet) ToJson

```go
func (s *Sheet) ToJson() datatype.JsonArray
```

ToJson

将 Sheet 转为 json Array 格式

## <a id="type-value">type</a> Value

```go
type Value struct {
    // 字符串字面量
    Liter string

    // 实际可解析类型
    Type ValueType

    // 真实解析后的值
    TrueValue any
}

```

Value
单元格元素数据内容

### <a id="func-newvalue">func</a> NewValue

```go
func NewValue(liter string) *Value
```

NewValue

创建 Value 结构体

### <a id="func-value-asfloat">func</a> (\*Value) AsFloat

```go
func (v *Value) AsFloat() (float64, error)
```

### <a id="func-value-asint">func</a> (\*Value) AsInt

```go
func (v *Value) AsInt() (int, error)
```

### <a id="func-value-asstring">func</a> (\*Value) AsString

```go
func (v *Value) AsString() (string, error)
```

### <a id="func-value-generrvalueconvert">func</a> (\*Value) GenErrValueConvert

```go
func (v *Value) GenErrValueConvert(intoType ValueType) *ErrValueConvert
```

GenErrValueConvert 生成类型转换错误

## <a id="type-valuetype">type</a> ValueType

```go
type ValueType uint
```

ValueType

单元格元素数据可解析的值

```go
const (
    // ValueTypeInt
    // int 类型
    ValueTypeInt ValueType = iota
    // ValueTypeFloat
    // float 类型
    ValueTypeFloat ValueType = iota
    // ValueTypeString
    // string 类型
    ValueTypeString ValueType = iota
    // ValueTypeNull
    // 自定义 Null 类型
    ValueTypeNull ValueType = iota
)
```

---
