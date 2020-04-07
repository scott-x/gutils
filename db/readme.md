# package db

### API
- `func GetConnection(config *model.DBConfig) (*sql.DB, error) `: get a database connection

### model
```golang
type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}
``` 
### example
```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/db"
	"github.com/scott-x/gutils/model"
)

func main() {
	//config
	config := &model.DBConfig{}
	//get connection, we'd better set it as a global value if it was used constantly.
	con, err := db.GetConnection(config)
	if err != nil {
		return
	}
	//define sql
	sql := "select username from user"

	//prepare
	stmt, err := con.Prepare(sql)
	if err != nil {
		panic(err)
	}
	//query
	row, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	for row.Next() {
		var username string
		row.Scan(&username)
		// handle result
		fmt.Println(username)
	}
}
```