package fs

import (
	"log"
	"os"
)

func Dir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// -1 error occurs
// 0 file
// 1 folder
func CheckFileType(file string) int {
	var res int = -1
	fi, err := os.Stat(file)
	if err != nil {
		// log.Fatal(err)
		return res
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		res = 1
	case mode.IsRegular():
		// do file stuff
		res = 0
	}
	return res
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists as a boolean
func DirExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}
