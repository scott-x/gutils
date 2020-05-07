package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

//get the size with 2 decimals
func GetFileSize(filename string) string {
	var result float64
	var size string
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = float64(f.Size())
		return nil
	})
	if result < 1024*1024 {
		f_2 := Decimal(result / 1024)
		size = float64_to_string(f_2) + "KB"
	} else if result < 1024*1024*1024 {
		f_2 := Decimal(result / (1024 * 1024))
		size = float64_to_string(f_2) + "MB"
	} else {
		f_2 := Decimal(result / (1024 * 1024 * 1024))
		size = float64_to_string(f_2) + "GB"
	}
	return size
}

//https://blog.csdn.net/wslyk606/article/details/81333001
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//https://yourbasic.org/golang/convert-string-to-float/
func float64_to_string(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
