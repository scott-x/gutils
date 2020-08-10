# package db

### API
- `func GetConnection(file string) (*sql.DB, error)`: get a database connection, if `err!=nil`, `*sql.DB` is nil, which means the connection can't be used.

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

config.json

```json
{
	"driver":"mysql",
	"host":"127.0.0.1",
	"username":"",
	"password":"",
	"port":"3306",
	"database":"test"
}
```
### example
```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/db"
	"log"
)

func main() {
	dbCon, err := db.GetConnection("db/config.json")
	if err != nil {
		log.Printf("DB error: %s\n", err)
		return
	}
	//define sql
	_sql := "select username from user"

	//prepare
	stmt, err := dbCon.Prepare(_sql)
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