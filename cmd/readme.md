## useage

## color option

```
number              color
------------------------------------
1                   red
2                   blue
3                   magenta
4                   yellow
5                   white
6                   cyan
7                   green
```

## API

- `func AddTask(tip string, color int, tasks ...string) string`: print the tasks and return the option you selected. If you pass `""` to tip, it will use build-in tip, otherwise it will use customed tip; color is a int number, which ranges from 1-7, default 6.
- `func AddQuestion(name, tip, retip, re string) *model.Questions`
- `func Exec() map[string]string`: return the result with map
- `func AskQuestion(tip string) string `
- `func Info(str string)`: print info
- `func Warning(str string)`: print warning info
- `func Trim(value string) string`: trim space of the value received from terminal
- `func SelectOne(desc, tip string, color_option int, t model.Tasker) int`: select one item from slice

## interface

```go
type Tasker interface {
	HandleItems() []string
}
```

```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/cmd"
)

func main() {
	// option := cmd.AddTask("swmiming", "eating", "sleeping") detatched
	option := cmd.AddTask("", 7, "swmiming", "eating", "sleeping")
	switch option {
	case "1":
		//do something
		task1()
	//anycode here ...
	case "2":
		//do something
	default:
		//do something
	}
}

func task1() {
	cmd.AddQuestion("name", "What's your name ? ", "Please input correct name: ", "^[a-z]+")
	cmd.AddQuestion("age", "What's your age ? ", "Please input correct age: ", "^[0-9]{2}$")
	answers := cmd.Exec()
	fmt.Println(answers)
}
```

