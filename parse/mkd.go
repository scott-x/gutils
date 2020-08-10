/*
* @Author: scottxiong
* @Date:   2020-04-20 14:29:30
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-08-10 15:51:09
 */
package parse

import (
	"github.com/russross/blackfriday/v2"
)

func MKD(mkdtext string) string {
	input := []byte(mkdtext)
	output := blackfriday.Run(input)
	return string(output)
}
