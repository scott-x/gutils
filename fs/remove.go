/*
* @Author: scottxiong
* @Date:   2019-07-25 18:45:44
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-06-10 16:41:30
 */
//https://learngolang.net/tutorials/how-to-remove-all-files-in-a-directory-in-go/
package fs

import (
	"log"
	"os"
)

func RemoveAll(folder string) {
	err := os.RemoveAll(folder)
	if err != nil {
		log.Printf("Delete Error: %s", err)
	}
}

func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Printf("Delete Error: %s", err)
	}
}
