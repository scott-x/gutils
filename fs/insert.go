/*
* @Author: scottxiong
* @Date:   2019-12-27 22:27:30
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-28 00:53:27
 */
package fs

import (
	"github.com/scott-x/gutils/model"
	"strings"
)

//insert the sentence after
func InsertAfter(insert *model.Insert) {
	new_arr := make([]string, 0)
	content, err := ReadFile1(insert.File)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(content, "\n")

	for k, v := range arr {
		new_arr = append(new_arr, v)
		if insert.Keywords != "" {
			if strings.Contains(v, insert.Keywords) {
				new_arr = append(new_arr, insert.NewLine)
			}
		}
		if insert.Line > 0 && insert.Line-1 == k && insert.Line-1 <= len(arr) {
			new_arr = append(new_arr, insert.NewLine)
		}

	}
	new_str := strings.Join(new_arr, "\n")
	if insert.Old != "" {
		new_str = strings.Replace(new_str, insert.Old, insert.New, -1)
	}
	//if Line, Keywords or Replace not given, do nothing
	if len(new_str) != len(content) {
		WriteString(insert.File, new_str)
	}
}

//insert the sentence before
func InsertBefore(insert *model.Insert) {
	new_arr := make([]string, 0)
	content, err := ReadFile1(insert.File)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(content, "\n")

	for k, v := range arr {
		if insert.Keywords != "" {
			if strings.Contains(v, insert.Keywords) {
				new_arr = append(new_arr, insert.NewLine)
			}
		}
		if insert.Line > 0 && insert.Line-1 == k && insert.Line-1 <= len(arr) {
			new_arr = append(new_arr, insert.NewLine)
		}
		new_arr = append(new_arr, v)

	}
	new_str := strings.Join(new_arr, "\n")
	if insert.Old != "" {
		new_str = strings.Replace(new_str, insert.Old, insert.New, -1)
	}
	//if Line, Keywords or Replace not given, do nothing
	if len(new_str) != len(content) {
		WriteString(insert.File, new_str)
	}
}
