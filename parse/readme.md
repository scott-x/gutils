# package parse

### parse sql

```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/parse"
)

func main() {
	tbs := parse.GetTables("test.sql")
	for _, table := range *tbs {
		fmt.Println("table:" + table.Name)
		for _, field := range table.Fields {
			fmt.Println(field.Name + ":" + field.Type)
		}
		fmt.Println("***********************")
	}
}
```
### example

```
table:user
id:int
username:string
password:string
public:int
see:int
create_time:string
***********************
table:video_info
id:string
author_id:int
name:string
display_ctime:string
public:int
see:int
create_time:string
***********************
table:comment
id:string
video_id:string
author_id:int
content:string
public:int
c_time:string
see:int
***********************
```