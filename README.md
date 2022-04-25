# XS

xs 全称是 xlsx server，这是一个简单易用的服务器，通过解析多个 xlsx 文件，建立多个 CRUD 服务。

通过 RESTful 风格的接口，可以实现对 xlsx 文件的增删改查，同时及时反映到服务器中。

# 使用方式

## 读取两个文件

```bash
xs -f dir1/example1.xlsx+dir1/example2.xlsx
```

## 读取目录

```bash
xs -f dir1/*.xlsx
```

## 读取多个目录和文件

```bash
xs -f dir1/*.xlsx+dir2/*.xlsx+dir3/example.xlsx
```

## 设置端口

```bash
xs -f *.xlsx -p 8000
```

## 读取文件

```bash
xs -c config.json
```

配置文件格式如下：

```json
{
	"port": 4040,
	"files": [
		{
			"file": "./testdata/xlsx/参加复试总名单.xlsx"
		},
		{
			"prefix": "test",
			"file": "./testdata/xlsx/testFile1.xlsx"
		}
	]
}
```

其中配置文件不支持通配符，但可以自定义 prefix 字段

# Export Packages

- [gitee.com/feimos/xs/cmd](./cmd/README.md)

- [gitee.com/feimos/xs/config](./config/README.md)

- [gitee.com/feimos/xs/controller](./controller/README.md)

- [gitee.com/feimos/xs/datatype](./datatype/README.md)

- [gitee.com/feimos/xs/server](./server/README.md)

- [gitee.com/feimos/xs/templates](./templates/README.md)

- [gitee.com/feimos/xs/utils](./utils/README.md)

- [gitee.com/feimos/xs/xlsx](./xlsx/README.md)
