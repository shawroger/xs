package templates

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"path"
)

const staticDir = "assets"

//go:embed assets
var assetsDir embed.FS

//go:embed index.html
var indexHtml string

type StaticFile struct {
	Content string
	Url     string
}

//ParseStaticFiles
//
//解析静态目录
func ParseStaticFiles() []StaticFile {

	var files []StaticFile
	templates, _ := fs.ReadDir(assetsDir, staticDir)

	for _, template := range templates {
		if !template.IsDir() {
			name := template.Name()
			url := path.Join(staticDir, name)

			content, err := assetsDir.ReadFile(url)
			if err != nil {
				panic(err)
			}

			staticFile := StaticFile{
				Url:     url,
				Content: string(content),
			}

			files = append(files, staticFile)
		}

	}

	return files
}

// BindIndexHtml
//
// 绑定 HTML 数据
func BindIndexHtml(params ...any) string {
	return fmt.Sprintf(indexHtml, params...)
}
