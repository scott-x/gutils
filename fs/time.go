/*
* @Author: scottxiong
* @Date:   2021-02-02 15:55:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2021-02-02 16:08:57
 */
package fs

import (
	"log"
	"os"
	"strconv"
	"time"
)

//获取文件修改时间 返回unix时间戳
func getFileModTime(pth string) int64 {
	f, err := os.Open(pth)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

//获取文件上此更改时候，可以用来判断文件是否修改
//比如获取excel信息，如果文件没改动就直接从database里获取，如果改动就直接从excel中获取
func LastModify(pth string) string {
	return strconv.FormatInt(getFileModTime(pth), 10)
}
