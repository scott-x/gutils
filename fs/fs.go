/*
* @Author: scottxiong
* @Date:   2019-07-25 16:10:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 01:47:04
 */
package fs

import (
	"github.com/otiai10/copy"
	"io"
	"log"
	"os"
)

func CopyFolder(src, des string) error {
	return copy.Copy(src, des)
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func Rename(oldName, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

//linux's pwd
func GetRunningFolder() string {
	str, _ := os.Getwd()
	return str
}

func NewFile(filePath string) {
	f, err := os.OpenFile(filePath, os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
