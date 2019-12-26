## useage

## API
- `func AddTask(tip string, color int, tasks ...string) string`: print the tasks and return the option you selected. If you pass `""` to tip, it will use build-in tip, otherwise it will use customed tip; color is a int number, which ranges from 1-6, default 6.

```
1 BoldRed
2 BoldBlue
3 BoldMagenta
4 BoldYellow
5 BoldWhite
6 BoldCyan
```

- `func AddQuestion(name, tip, retip, re string) *model.Questions`
- `func Exec() map[string]string`: return the result with map
- `func AskQuestion(tip string) string `

```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/cmd"
)

func main() {

	option := cmd.AddTask("swmiming", "eating", "sleeping")
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

