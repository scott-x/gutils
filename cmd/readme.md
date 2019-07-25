## useage

```golang
package main

import (
	"github.com/scott-x/gutils/cmd"
	"fmt"
)
func main(){
   cmd.AddQuestion("name","What's your name ? ","Please input correct name: ","[a-z]+")
   cmd.AddQuestion("age","What's your age ? ","Please input correct age: ","[0-9]{2}")
   answers :=cmd.Exec()
   fmt.Println(answers) 
   //anycode here ...
}   
```