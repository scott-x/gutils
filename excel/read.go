/*
* @Author: scottxiong
* @Date:   2020-01-14 17:41:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-01-15 09:14:30
 */
package excel

import (
	_ "fmt"
	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
	_ "strings"
)

//excel
type XLS struct {
	Name string
}

type XLSX struct {
	Name string
}

type Position struct {
	Sheet_index int //starts from 0
	Row         int //row index 1->0 2->1 ...
	Col         int //row index A->0 B->1 if it has multiple columns, the index value depends on the first column, equal to first_column-1
}

type Positions []Position

type Read interface {
	read(p *Position) string
}

type Reads interface {
	reads(ps *Positions) []string
}

func (x *XLS) reads(ps *Positions) []string {
	xlsFile, err := xls.Open(x.Name, "utf-8")
	if err != nil {
		panic(err)
	}
	res := make([]string, 0)
	for _, p := range *ps {
		res = append(res, xlsFile.GetSheet(p.Sheet_index).Row(p.Row).Col(p.Col))
	}
	return res
}

func (xs *XLSX) reads(ps *Positions) []string {
	xlFile, err := xlsx.OpenFile(xs.Name)
	if err != nil {
		panic(err)
	}
	res := make([]string, 0)
	for _, p := range *ps {
		res = append(res, xlFile.Sheets[p.Sheet_index].Rows[p.Row].Cells[p.Col].Value)
	}
	return res
}

func (x *XLS) read(p *Position) string {
	xlsFile, err := xls.Open(x.Name, "utf-8")
	if err != nil {
		panic(err)
	}
	return xlsFile.GetSheet(p.Sheet_index).Row(p.Row).Col(p.Col)
}

func (xs *XLSX) read(p *Position) string {
	xlFile, err := xlsx.OpenFile(xs.Name)
	if err != nil {
		panic(err)
	}
	return xlFile.Sheets[p.Sheet_index].Rows[p.Row].Cells[p.Col].Value
}

func GetValue(r Read, p *Position) string {
	return r.read(p)
}

func GetValues(r Reads, ps *Positions) []string {
	return r.reads(ps)
}
