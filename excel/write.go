/*
* @Author: scottxiong
* @Date:   2020-01-15 10:12:42
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-01-15 11:06:04
 */
package excel

import (
	"github.com/tealeg/xlsx"
)

type WriteInfo struct {
	FileName  string // eg: a.xlsx
	SheetName string // eg: 任务单
	Row       int
	Col       int
	Value     string
}

func Write(w *WriteInfo) {
	file, error := xlsx.OpenFile(w.FileName)
	defer file.Save(w.FileName)
	if error != nil {
		panic(error)
	}
	file.Sheet[w.SheetName].Rows[w.Row].Cells[w.Col].SetString(w.Value)
}
