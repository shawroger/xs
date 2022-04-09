package xlsx

import "github.com/xuri/excelize/v2"

// Object
//
// 对于每一个 XLSX 文件
//
// 生成一个独立的 Object 对象
type Object struct {
	Sheets *[]Sheet
}

// NewObjectFromFile
//
// 从 *excelize.File 句柄生成 Object 对象
func NewObjectFromFile(f *excelize.File) (*Object, error) {
	var sheets []Sheet
	sheetList := f.GetSheetList()
	for _, sheetName := range sheetList {
		var sheet Sheet
		sheet.Name = sheetName

		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, err
		}

		for line, row := range rows {
			// 第一行
			// 储存键名
			if line == 0 {
				var keys []string
				for _, key := range row {
					keys = append(keys, key)
				}
				sheet.Key = keys
			} else {
				parsedRow := ParseRowFromStringArr(sheet.Key, row)
				sheet.Rows = append(sheet.Rows, *parsedRow)
			}
		}
		sheets = append(sheets, sheet)
	}

	return &Object{Sheets: &sheets}, nil
}
