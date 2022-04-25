# cmd

`import "gitee.com/feimos/xs/cmd"`

- [Overview](#overview)
- [Export List](#export-list)

## <a id="overview">Overview</a>

## <a id="export-list">Export List</a>

- [func AppInit(a *datatype.AppInfo) *cli.App](#func-appinit)
- [type ErrMissingDataSource](#type-errmissingdatasource)
  - [func (s ErrMissingDataSource) Error() string](#func-errmissingdatasource-error)
- [type Flags](#type-flags)

## <a id="func-appinit">func</a> AppInit

```go
func AppInit(a *datatype.AppInfo) *cli.App
```

## <a id="type-errmissingdatasource">type</a> ErrMissingDataSource

```go
type ErrMissingDataSource struct {
}

```

### <a id="func-errmissingdatasource-error">func</a> (ErrMissingDataSource) Error

```go
func (s ErrMissingDataSource) Error() string
```

## <a id="type-flags">type</a> Flags

```go
type Flags struct {
    // contains filtered or unexported fields
}

```

Flags

cmd 的 flag 信息

---
