# package excel
### API

- `func GetValue(r Read, p *Position) string`
- `func GetValues(r Reads, ps *Positions) []string`
- `func Write(w *WriteInfo)`

```
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

type WriteInfo struct {
	FileName  string // eg: a.xlsx
	SheetName string // eg: 任务单
	Row       int
	Col       int
	Value     string
}
```
### example write

xlxs
```golang
package main

import (
	"github.com/scott-x/gutils/excel"
)

func main() {
	info := &excel.WriteInfo{
		"a.xlsx",
		"Sheet1",
		1,//row index
		1, //
		"hello",
	}
	excel.Write(info)
}

```

### example read

get a single value
```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/excel"
)

func main() {
	r := &excel.XLS{"a.xls"}
	p := &excel.Position{0, 1, 1}
	fmt.Println(excel.GetValue(r, p))
}
```

get mutiple value

```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/excel"
)

func main() {
	r := &excel.XLSX{"a.xlsx"}
	ps := &excel.Positions{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
	}
	fmt.Println(excel.GetValues(r, ps))
}
```