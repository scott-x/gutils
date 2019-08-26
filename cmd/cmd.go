/*
* @Author: sottxiong
* @Date:   2019-07-07 16:28:34
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-27 00:08:58
 */
package cmd

import (
	"bufio"
	"github.com/fatih/color"
	"github.com/scott-x/gutils/cl"
	"os"
	"regexp"
	"strings"
)

type Question struct {
	Name  string
	Tip   string
	ReTip string
	Re    string
	//Do func()
}

type Questions struct {
	qs []Question
}

var (
	questions = &Questions{}
	answers   = map[string]string{}
	all       = 0
	cmded     = 0
)

func AddQuestion(name, tip, retip, re string) *Questions {
	all++
	questions.qs = append(questions.qs, Question{name, tip, retip, re})
	return questions
}

func ask_question(q string, color *color.Color) string {
	inputReader := bufio.NewReader(os.Stdin)

	color.Printf(q)
	inputData, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.Trim(inputData, "\n")

}

func command(q *Question) {
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

func Exec() map[string]string {
	//Exec()第二次执行的时候会出现bug

	for _, question := range questions.qs[cmded:all] {
		command(&question)
	}
	return answers
}
