/*
* @Author: scottxiong
* @Date:   2019-12-28 02:46:22
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-28 03:01:05
 */
package main

import (
	"fmt"
	"github.com/scott-x/gutils/fs"
)

func main() {
	a := fs.ListFiles("./")
	fmt.Println(a)
}
