/*
* @Author: scottxiong
* @Date:   2019-07-25 16:15:00
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 18:47:45
 */
package fs

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	CopyFile("a.txt", "b.txt")
	fmt.Println("Copy done!")
	// Rename("a.txt", "a.b.txt")
	// WriteString("a.txt", "hello world")
	// re := make(map[string]string)
	// re["am"] = "AM"
	// re["I"] = "i love"
	// CopyAndReplace("a.txt", "c.txt", re)
	//CopyFolder("/Users/apple/go/src/github.com/scott-x/gutils/", "/Users/apple/desktop/a")
	//CreateDirIfNotExist("/Users/apple/desktop/a/b/c")
	RemoveAll("/Users/apple/desktop/tes")
}
