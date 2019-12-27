/*
* @Author: scottxiong
* @Date:   2019-07-25 16:10:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-28 03:01:36
ref https://stackoverflow.com/questions/8824571/golang-determining-whether-file-points-to-file-or-directory
*/
package fs

import (
	"github.com/otiai10/copy"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func List(folder string) []string {
	f := make([]string, 0)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f = append(f, file.Name())
	}
	return f
}

func ListFolder(folder string) []string {
	f := make([]string, 0)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		t := CheckFileType(file.Name())
		if t == 1 {
			f = append(f, file.Name())
		}
	}
	return f
}

// just list current folder, if folder, will be ignored
func ListFiles(folder string) []string {
	f := make([]string, 0)
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		t := CheckFileType(file.Name())
		if t == 0 {
			f = append(f, file.Name())
		}
	}
	return f
}

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
