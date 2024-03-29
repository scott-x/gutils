/*
* @Author: scottxiong
* @Date:   2019-07-25 17:14:40
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-08 07:13:30
 */
package fs

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile1(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadAndReplace(file string, replace map[string]string) error {
	content, err := ReadFile1(file)
	if err != nil {
		return err
	}
	var new_content string = string(content)
	for k, v := range replace {
		new_content = strings.ReplaceAll(string(new_content), k, v)
	}
	WriteString(file, new_content)
	return nil
}

func ReadFile(file string) string {
	/*
	   先用os的Open方法获取文件句柄
	*/
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	//程序退出时关闭文件（释放资源）
	defer f.Close()

	inputReader := bufio.NewReader(f)
	var data string = ""
	//在一个无限循环中使用 ReadString('\n') 或 ReadBytes('\n') 将文件的内容逐行（行结束符 '\n'）读取出来。
	for {
		inputString, readerError := inputReader.ReadString('\n')
		//fmt.Printf("The input was: %s", inputString)
		data += inputString
		/*一旦读取到文件末尾，变量 readerError 的值将变成非空（事实上，常量 io.EOF 的值是 true），
		我们就会执行 return 语句从而退出循环。*/
		if readerError == io.EOF {
			return data
		}
	}

}

//将整个文件的内容读到一个字符串里：
/*
io/ioutil 包里的 ioutil.ReadFile() 方法，该方法第一个返回值的类型是 []byte，
里面存放读取到的内容，第二个返回值是错误，如果没有错误发生，第二个返回值为 nil
*/
func Copy(inputFile, outputFile string) {
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	//fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

func CopyAndReplace(inputFile, outputFile string, replace map[string]string) {
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	//fmt.Println(string(buf))
	var newStr string = ""
	for k, v := range replace {
		newStr = strings.ReplaceAll(string(buf), k, v)
	}
	err = ioutil.WriteFile(outputFile, []byte(newStr), 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
