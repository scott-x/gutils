/*
除了文件句柄，我们还需要 bufio 的 Writer。我们以只写模式打开文件 output.dat，如果文件不存在则自动创建：
outputFile, outputError := os.OpenFile(“output.dat”, os.O_WRONLY|os.O_CREATE, 0666)
可以看到，OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符“|”连接），使用的文件权限。
我们通常会用到以下标志：
os.O_RDONLY：只读
os.O_WRONLY：只写
os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。
在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用0。而在写文件时，
不管是 Unix 还是 Windows，都需要使用 0666。

ref :https://golangbot.com/write-files/
*/

package fs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteBuf(file, data string) {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := data
	outputWriter.WriteString(outputString)
	outputWriter.Flush()
}

// func WriteString(file, data string) {
// 	f, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0666) // 要想写入数据，许给0666的写入权限
// 	defer f.Close()
// 	f.WriteString(data)
// }

func WriteString(file, data string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func IsExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
