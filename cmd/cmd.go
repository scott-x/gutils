/*
* @Author: sottxiong
* @Date:   2019-07-07 16:28:34
* @Last Modified by:   scottxiong
* @Last Modified time: 2022-01-12 21:42:40
 */
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/model"
)

var (
	questions = &model.Questions{}
	answers   = map[string]string{}
	all       = 0
	cmded     = 0
	isWindows bool
)

func init() {
	if runtime.GOOS == "windows" {
		isWindows = true
	}
}

func AddQuestion(name, tip, retip, re string) *model.Questions {
	all++
	questions.Qs = append(questions.Qs, model.Question{name, tip, retip, re})
	return questions
}

func AskQuestion(tip string) string {
	q := &model.SimpleQuestion{
		Tip: tip,
	}
	return ask_question(q.Tip, cl.BoldGreen)
}

func ask_question(q string, color *color.Color) string {
	var inputData string
	var err error
	if isWindows {
		_, err = fmt.Scanf("%s", &inputData)
		if err != nil {
			return ""
		}
	} else {
		inputReader := bufio.NewReader(os.Stdin)

		color.Printf(q)
		inputData, err = inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		inputData = strings.Trim(inputData, "\n")
	}
	return inputData
}

func command(q *model.Question) {
	data := ask_question(q.Tip, cl.BoldGreen)
	re := regexp.MustCompile(q.Re)
	for {
		match := re.FindString(data)
		if len(match) > 0 {
			answers[q.Name] = data
			break
		} else {
			data = ask_question(q.ReTip, cl.BoldRed)
		}
	}
	//q.Do()
	cmded++
}

func AddTask(tip string, color_option int, tasks ...string) string {
	ts := model.Tasks{}
	for _, t := range tasks {
		ts.Names = append(ts.Names, t)
	}
	tasks_length := len(ts.Names)
	//color selection
	var my_color *color.Color
	switch color_option {
	case 1:
		my_color = cl.BoldRed
	case 2:
		my_color = cl.BoldBlue
	case 3:
		my_color = cl.BoldMagenta
	case 4:
		my_color = cl.BoldYellow
	case 5:
		my_color = cl.BoldWhite
	case 6:
		my_color = cl.BoldCyan
	case 7:
		my_color = cl.BoldGreen
	default:
		my_color = cl.BoldCyan
	}
	//use default tips
	if tip == "" {
		if tasks_length == 1 {
			my_color.Printf("What you expected to do?\n")
		} else {
			my_color.Printf("What you expected to do? Select the number 1-%d: \n", tasks_length)
		}
	} else {
		//customed tip
		my_color.Printf("%s\n", tip)
	}

	for k, t := range ts.Names {
		cl.BoldCyan.Printf(" %d. %s\n", k+1, t)
	}
	cl.BoldRed.Printf("My selection is: ")
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	option = strings.Trim(option, "\n")
	re := `^[1-9][0-9]*$`
	res := regexp.MustCompile(re)
	str := res.FindAllString(option, -1)

	if len(str) == 0 {
		if tasks_length == 1 {
			cl.BoldRed.Printf("Please input number 1 \n")
			return ""
		} else {
			cl.BoldRed.Printf("Please input number from 1 to %d \n", tasks_length)
			return ""
		}
	}
	num, _ := strconv.Atoi(option)
	if num > 0 && num <= tasks_length {
		return option
	} else {
		cl.BoldRed.Printf("Please input number from 1 to %d \n", tasks_length)
		return ""
	}
}

//return selected num, if invalid return -1
//desc: description
func SelectOne(desc, tip string, color_option int, t model.Tasker) int {
	display := t.HandleItems()
	tasks_length := len(display)
	//color selection
	var my_color *color.Color
	switch color_option {
	case 1:
		my_color = cl.BoldRed
	case 2:
		my_color = cl.BoldBlue
	case 3:
		my_color = cl.BoldMagenta
	case 4:
		my_color = cl.BoldYellow
	case 5:
		my_color = cl.BoldWhite
	case 6:
		my_color = cl.BoldCyan
	case 7:
		my_color = cl.BoldGreen
	default:
		my_color = cl.BoldCyan
	}

	//desc
	if tasks_length == 1 {
		my_color.Printf("%s:\n", desc)
	} else {
		my_color.Printf("%s, please select number [1-%d]:\n", desc, tasks_length)
	}

	//use default tips
	if tip == "" {
		if tasks_length == 1 {
			my_color.Printf("What you expected to do?\n")
		} else {
			my_color.Printf("What you expected to do? Select the number 1-%d: \n", tasks_length)
		}
	} else {
		//customed tip
		my_color.Printf("%s\n", tip)
	}

	for k, t := range display {
		cl.BoldCyan.Printf(" %d. %s\n", k+1, t)
	}
	cl.BoldRed.Printf("My selection is: ")
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	option = strings.Trim(option, "\n")
	re := `^[1-9][0-9]*$`
	res := regexp.MustCompile(re)
	str := res.FindAllString(option, -1)

	if len(str) == 0 {
		if tasks_length == 1 {
			cl.BoldRed.Printf("Please input number 1 \n")
			return -1
		} else {
			cl.BoldRed.Printf("Please input number from 1 to %d \n", tasks_length)
			return -1
		}
	}
	num, _ := strconv.Atoi(option)
	if num > 0 && num <= tasks_length {
		return num
	} else {
		cl.BoldRed.Printf("Please input number from 1 to %d \n", tasks_length)
		return -1
	}
}

func Exec() map[string]string {
	//Exec()第二次执行的时候会出现bug

	for _, question := range questions.Qs[cmded:all] {
		command(&question)
	}
	return answers
}

func Trim(value string) string {
	return strings.TrimSpace(value)
}
