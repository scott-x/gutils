/*
* @Author: sottxiong
* @Date:   2019-12-29 22:05:59
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-12-29 22:10:53
*/
package cmd

import (
	"github.com/scott-x/gutils/cl"
)

func Info(str string){
	cl.BoldCyan.Printf("info: ")
	cl.BoldGreen.Printf("%s\n",str)
}

func Warning(str string) {
	cl.BoldRed.Printf("warning: ")
	cl.BoldGreen.Printf("%s\n",str)
}