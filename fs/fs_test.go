/*
* @Author: scottxiong
* @Date:   2019-07-25 16:15:00
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 16:17:40
 */
package fs

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	CopyFile("a.txt", "b.txt")
	fmt.Println("Copy done!")
}
