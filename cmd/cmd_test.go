/*
* @Author: scottxiong
* @Date:   2019-07-25 15:59:01
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 16:04:34
 */
package cmd

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	AddQuestion("name", "What's your name ? ", "Please input correct name: ", `^[a-z]+`)
	AddQuestion("age", "What's your age ? ", "Please input correct age: ", `^[0-9]{2}$`)
	answers := Exec()
	fmt.Println(answers)
	//anycode here ...
}
