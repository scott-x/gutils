/*
* @Author: scottxiong
* @Date:   2019-08-27 00:03:50
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-27 00:03:55
 */

package main

import (
	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/cmd"
)

func main() {
	cl.BoldYellow.Printf("hello, my name is Scott \n")
	cmd.AddQuestion("name", "What's your name ? ", "Please input correct name: ", "^[a-z]+")
	cmd.AddQuestion("age", "What's your age ? ", "Please input correct age: ", "^[0-9]{2}$")
	answers := cmd.Exec()
	cl.BoldRed.Println(answers)
	cmd.AddQuestion("a", "What's a? ", "Please input correct name: ", "^[a-z]+")
	cmd.AddQuestion("b", "What'sb ? ", "Please input correct age: ", "^[0-9]{2}$")
	answers = cmd.Exec()
	cl.BoldRed.Println(answers)
}
