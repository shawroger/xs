# config

`import "gitee.com/feimos/xs/config"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [type Config](#type-config)
- [func DefaultConfig() \*Config](#func-defaultconfig)
- [func NewConfigFromSingleFile(filename ...string) \*Config](#func-newconfigfromsinglefile)
- [func ParseFile(filename string) (\*Config, error)](#func-parsefile)
- [type FileConfig](#type-fileconfig)
- [func CreateDefaultFromFilename(file string) FileConfig](#func-createdefaultfromfilename)

## <a id="type-config">type</a> Config

```go
type Config struct {
    Port     int
    Debug    bool
    GinDebug bool
    Files    []FileConfig
}

```

Config

配置项 结构体

### <a id="func-defaultconfig">func</a> DefaultConfig

```go
func DefaultConfig() *Config
```

### <a id="func-newconfigfromsinglefile">func</a> NewConfigFromSingleFile

```go
func NewConfigFromSingleFile(filename ...string) *Config
```

NewConfigFromSingleFile

从单一文件路径创建 config

### <a id="func-parsefile">func</a> ParseFile

```go
func ParseFile(filename string) (*Config, error)
```

ParseFile

解析文件，获得配置结构体

## <a id="type-fileconfig">type</a> FileConfig

```go
type FileConfig struct {
    // 文件路径
    File string
    // 使用的前缀
    Prefix string
    // 筛选工作表
    //
    // 为空，则是为全部开启
    Sheets []string
}

```

FileConfig

文件配置项

### <a id="func-createdefaultfromfilename">func</a> CreateDefaultFromFilename

```go
func CreateDefaultFromFilename(file string) FileConfig
```

CreateDefaultFromFilename

从单文件路径生成默认 FileConfig

---
