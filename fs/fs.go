/*
* @Author: scottxiong
* @Date:   2019-07-25 16:10:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-12 04:57:00
ref https://stackoverflow.com/questions/8824571/golang-determining-whether-file-points-to-file-or-directory
*/
package fs

import (
	"archive/zip"
	"github.com/otiai10/copy"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type F struct {
	Path string
	Size int64
}

type FS []F

type INFO struct {
	Path   string //file path
	Size   int64
	Header *zip.FileHeader
}

type INFOS []INFO

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

func ListAll(folder string, ignore []string) ([]string, error) {
	f := make([]string, 0)
	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			n := FileType(path)
			if n == 1 {
				//file
				if len(ignore) > 0 {
					for _, v := range ignore {
						if !strings.Contains(path, v) {
							f = append(f, path)
						}
					}
				}

			}
			// fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return f, nil
}

func ListAllWithFileHeaders(folders []string) (*INFOS, int64, error) {
	var err error
	var sum int64
	infos := &INFOS{}

	for _, folder := range folders {
		err = filepath.Walk(folder,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				n := FileType(path)

				if n == 1 {
					//file
					sum += info.Size()
					i := &INFO{}
					header, err := zip.FileInfoHeader(info)
					if err != nil {
						return err
					}
					i.Header = header
					i.Path = path
					i.Size = info.Size()
					*infos = append(*infos, *i)

				} else {

				}

				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
	return infos, sum, nil
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
