/*
* @Author: scottxiong
* @Date:   2019-07-25 15:05:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 15:29:03
 */
package html

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const url = "https://www.baidu.com"
	bytes, err := Fetch(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("html content: %s", bytes)
}
