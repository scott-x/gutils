/*
* @Author: scottxiong
* @Date:   2019-07-25 16:15:00
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 19:12:03
 */
package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	// CopyFile("a.txt", "b.txt")
	// fmt.Println("Copy done!")
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	WriteString("a.json", string(b))
	// data := ReadFile("a.json")
	// fmt.Println(data)
	// Rename("a.txt", "a.b.txt")
	// WriteString("a.txt", "hello world")
	// re := make(map[string]string)
	// re["am"] = "AM"
	// re["I"] = "i love"
	// CopyAndReplace("a.txt", "c.txt", re)
	//CopyFolder("/Users/apple/go/src/github.com/scott-x/gutils/", "/Users/apple/desktop/a")
	//CreateDirIfNotExist("/Users/apple/desktop/a/b/c")
	// RemoveAll("/Users/apple/desktop/tes")
}
