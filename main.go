/*
* @Author: scottxiong
* @Date:   2020-01-13 15:09:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-01-13 15:11:05
 */
package main

import (
	"fmt"
	"github.com/scott-x/gutils/fs"
)

func main() {
	fmt.Println(fs.ListFiles("./"))
}
