/*
* @Author: scottxiong
* @Date:   2020-01-14 17:41:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-01-14 18:02:03
 */
package excel

import (
	"fmt"
	"github.com/extrame/xls"
	"github.com/scott-x/gutils/model"
	"github.com/tealeg/xlsx"
	"strings"
)

type Read interface {
	read(p *model.Position) string
}

func (x *model.XLS) read(p *model.Position) string {
	xlsFile, err := xls.Open(x.Name, "utf-8")
	if err != nil {
		panic(err)
	}
	return xlsFile.GetSheet(p.sheet_index).Row(p.row).Col(p.col)
}

func (xs *model.XLSX) read(p *model.Position) string {
	xlFile, err := xlsx.OpenFile(xs.Name)
	if err != nil {
		panic(err)
	}
	return xlFile.Sheets[p.sheet_index].Rows[p.row].Cells[p.col].Value
}

func GetValue(r Read, p *model.Position) string {
	return r.read(p * model.Position)
}
