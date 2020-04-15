/*
* @Author: scottxiong
* @Date:   2019-12-27 22:27:30
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-15 16:22:53
 */
package fs

import (
	"errors"
	"fmt"
	"github.com/scott-x/gutils/model"
	"strings"
)

func Insert(insert *model.Insert) {
	if insert.Postion != model.POSITION_BEFORE && insert.Postion != model.POSITION_AFTER {
		fmt.Println(errors.New("Insert.Postion must be specified, model.POSTION_BEFORE or model.POSTION_AFTER will be OK"))
	}
	new_arr := make([]string, 0)
	content, err := ReadFile1(insert.File)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(content, "\n")

	for k, v := range arr {
		switch insert.Postion {
		case model.POSITION_AFTER: //insert the sentence after
			new_arr = append(new_arr, v)
			if insert.Keywords != "" {
				if strings.Contains(v, insert.Keywords) {
					new_arr = append(new_arr, insert.Content)
				}
			}
			if insert.Line > 0 && insert.Line-1 == k && insert.Line-1 <= len(arr) {
				new_arr = append(new_arr, insert.Content)
			}
		case model.POSITION_BEFORE: //insert the sentence before

			if insert.Keywords != "" {
				if strings.Contains(v, insert.Keywords) {
					new_arr = append(new_arr, insert.Content)
				}
			}
			if insert.Line > 0 && insert.Line-1 == k && insert.Line-1 <= len(arr) {
				new_arr = append(new_arr, insert.Content)
			}
			new_arr = append(new_arr, v)
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
