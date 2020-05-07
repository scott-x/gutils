/*
* @Author: scottxiong
* @Date:   2019-07-25 18:45:44
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-08 03:32:42
 */
//https://learngolang.net/tutorials/how-to-remove-all-files-in-a-directory-in-go/
package fs

import (
	"fmt"
	"os"
)

func RemoveAll(folder string) {
	err := os.RemoveAll(folder)
	fmt.Println(err)
}

func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
}
