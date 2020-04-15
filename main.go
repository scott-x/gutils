/*
* @Author: scottxiong
* @Date:   2020-04-15 19:03:53
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-15 19:18:19
 */
package main

import (
	"fmt"
	"github.com/scott-x/gutils/str"
)

func main() {
	sub := str.GetContentBetween("hello (wold(a,b) yds", " wold(", ")")
	fmt.Println(sub)
}
