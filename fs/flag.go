package fs

import (
	"fmt"
	"os"
)

// 0 dir, 1 file, -1 error
func FileType(filename string) int {
	var t int
	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		t = -1
		return t
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		t = 0
	case mode.IsRegular():
		// do file stuff
		t = 1
	}
	return t
}
