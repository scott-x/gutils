# package fs

### API
- `func Dir() string`: return the exec binary file folder
- `func RemoveAll(folder string)`: rm -rf
- `func RemoveFile(file string)`: remove one file
- `func Rename(oldName, newName string)`:rename the file or folder
- `func WriteBuf(file, data string) `: wirte a string data into the file
- `func WriteString(file, data string) `wirte a string data into the file
- `func Copy(from, to string) `: copy file, must specify what to it is.
- `func CopyFile(srcName, dstName string) (written int64, err error)`:copy file
- `func CopyAndReplace(inputFile, outputFile string, replace map[string]string)`:copy and replace
- `func CreateDirIfNotExist(dir string)`: similar like `mkdir -p /a/b/c`
- `func CopyFolder(src, des string) error`:copy folder, for example: `CopyFolder("/Users/apple/go/src/github.com/scott-x/gutils", "/Users/apple/desktop/a")` :copy all files from gutils to folder `a`, not include folder `gutils`
- `func IsExist(file string) bool`: check file or folder if exists
- `func ReadAndReplace(file string, replace map[string]string) error`: read file and replace it's content with `map[string]string{"old1":"new1","old2":"new2"}`
- `func ReadFile1(file string) (string, error)`: read the file and get the string content
- `func GetRunningFolder() string `: get running folder, same as linux's command `pwd`
- `func GetEnv(env string) string`:get the environment of the system
- `func NewFile(filePath string) `: create file
- `func InsertAfter(insert *model.Insert)`:insert after with replace function
- `func InsertBefore(insert *model.Insert)`:insert before with replace function
- `func CheckFileType(file string) int`: 0 indicates file, 1 indicates folder, -1 means doesn't exist.
- `func IsDirectory(path string) (bool, error) `: if path doesn't exist or path is a file, it will return false.
- `func FileType(filename string) int`:0 folder, 1 file, 0 error.
- `func List(folder string) []string `: list the folder and files in current folder, doesn't loop inner folder.
- `func ListFiles(folder string) []string`:just list the files in current folder
- `func ListFolder(folder string) []string`:only list folder in current folder
- `func ListAll(folder string, ignore []string) ([]string, error)`: list all, but will ignore the matched substring in the full path.
- `func ListAllWithFileHeaders(folders []string) (*INFOS, int64, error)`:
- `func Zip(zipName string, Base string, files []string) `: zip a file, `Base` will be removed.
- `func ZipWithBar(z *ZIP) `: zip with progress bar
- `func Tab(n int) string`: tabale n
- `func ReadJson(filename string) func(string) interface{}`: read the configuration of json, return a method, with which we can get the related value.
- `func ModifyAttrOfJson_STRING(filename string, key string, value string)`: modify the string value
- `func ModifyAttrOfJson_FLOAT64(filename string, key string, value float64)`: modify the data of float64
- `func MD5(file string) string`: md5 file checksum
- `func GetExpectedPath(folder,re string) string`: it will loop the folder, and return the file/folder that **1st** matches re, will return "" if not match.

### Attribute

- `fs.HOME`
- `fs.DESKTOP`

### import struct

```golang
//if Line, Keywords or Replace not given, do nothing
//insert structure
type Insert struct {
	File     string
	Content  string //new included content
	Postion  int    //before or after
	Line     int    //which line
	Keywords string //locate the line as per keywords that be given
	Replace         // optional, if you want to exec replace operation, add it on
}

//replace
type Replace struct {
	Old string
	New string
}

type F struct {
	Path string
	Size int64
}

type FS []F

//zip info
type ZIP struct {
	Folders []string
	Where   string
	Base    string // will be truncated from the full path
}

type INFO struct {
	Path   string //file path
	Size   int64
	Header *zip.FileHeader
}

type INFOS []INFO
```

### example

```golang
package main

import (
	_ "fmt"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/model"
)

func main() {
	insert := &model.Insert{
		File: "app.js",
		Content: `import { Hello } from "hello";

function add (a, b) {
	return a+b
}
`,
		Postion: model.POSITION_BEFORE,
		// Line
		Keywords: "ReactDOM.render",
		// Replace        :
	}

	fs.Insert(insert)
}
```

### zip

```golang
package main

import (
	"github.com/scott-x/gutils/fs"
)

func main() {
	folders := []string{
		"/Users/apple/Desktop/C2004F4_FLW/AI_ThisFolderToPrinter",
		"/Users/apple/Desktop/C2004F4_FLW/PDF_Locked_For_Visual_Ref",
	}
	zip := &fs.ZIP{
		Folders: folders,
		Where:   "/Users/apple/Desktop/C2004F4_FLW.zip",
		Base:    "/Users/apple/Desktop/C2004F4_FLW",
	}

	fs.ZipWithBar(zip)
}
```

### read configration

** 1. json **

config.json

```json
{
  "date": "2019-04-30",
  "mysql": {
    "url": "127.0.0.1:3306",
    "username": "root",
    "password": "123456"
  },
  "port" : 8888
}
```

```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/fs"
)

func main() {
	get := fs.ReadJson("config.json")
	urlValue := get("mysql.url").(string)
	port := get("port").(float64)
	fmt.Println("mysql url:", urlValue)
	fmt.Printf("mysql url: %s\nmysql username: %s\nmysql password: %s\n", get("mysql.url"), get("mysql.username"), get("mysql.password"))
}
```

### Walk

- `func NewTasks(pths []string, e *Expect) *[]Task`
- `func (t *Task) Walk(c chan string, done chan bool) error`

```golang
//exected file type
const (
	FILE = iota
	FOLDER
	MIX
)

type Task struct {
	Dir string
	*Expect
}

//constraint
type Expect struct {
	T             int // YOUR TARGET FILE TYPE, only FILE,FOLDER,MIX can be used here
	Match         string //regexp string, used to filter the data
	IgnoreFolders []string //ignore folders contains that defined the []string
	Depth         int 
	/* If your file type is folder, then you can set as 1, when found the target folder, 
	it won't walk the target folder; if not set, will walk all folder*/
}
```

**Example1: walk single folder**

```golang
package main


import (
	"github.com/scott-x/gutils/fs"
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	done := make(chan bool)
	num,taskHaveDone := 0,0

	pts := []string{
		"/Users/scottxiong/go",
		"/Users/scottxiong/Desktop",
	}

	e := &fs.Expect{
		T:     fs.FILE,
		Match: ".go$",
		Depth: -1,
		IgnoreFolders: []string{
			"test",
		},
	}
	len_pts := len(pts)
	
	t1 := time.Now()
	tasks := fs.NewTasks(pts, e)
	for _, f := range  *tasks{
		_f := f
		go _f.Walk(c, done)
	}

	for {
		select {
		case v := <-c:
			num++
			fmt.Println(v)
		case <-done:
			taskHaveDone++
			if taskHaveDone == len_pts{
				fmt.Println(num)
				fmt.Println(time.Since(t1))
				return
			}
		}
	}
}
```

**Example1: walk multiple folder**

```golang
func main() {
	c := make(chan string)   //design chan
	c1 := make(chan string)  //pub chan
	xls := make(chan string) //xls
	done := make(chan bool)
	num, taskHaveDone := 0, 0

	pts := []string{
		"/Volumes/datavolumn_bmkserver_Design/Proofing",
		"/Volumes/datavolumn_bmkserver_Design/WMT-Canada",
		"/Volumes/datavolumn_bmkserver_Design/WMT-USA",
	}

	pts1 := []string{
		"/Volumes/datavolumn_bmkserver_Pub/新做稿/未开始",
		"/Volumes/datavolumn_bmkserver_Pub/新做稿/进行中",
		"/Volumes/datavolumn_bmkserver_Pub/新做稿/已结束/WMT_Case",
		"/Volumes/datavolumn_bmkserver_Pub/新做稿/已结束/NON-WMT",
	}

	pts2 := []string{
		"/Volumes/datavolumn_bmkserver_Pub/新做稿/印刷",
	}

	total_len := len(pts) + len(pts1) + len(pts2)

	e := &fs.Expect{
		T:     FOLDER,
		Match: "^[UCBP][2][01][01][0-9][A-Z0-9][0-9]_[A-Z]{3}$",
		Depth: 1,
		IgnoreFolders: []string{
			"2019",
			"2020",
		},
	}

	e1 := &fs.Expect{
		T:     FOLDER,
		Match: "^[UCB][2][01][01][0-9][A-Z0-9][0-9]_[A-Z]{3} 做稿$",
		Depth: 1,
		IgnoreFolders: []string{
			"2019",
			"2020",
		},
	}

	e2 := &fs.Expect{
		T:     FOLDER,
		Match: "^[P][2][01][01][0-9][A-Z0-9][0-9]_LNC$",
		Depth: 1,
		IgnoreFolders: []string{
			"2018",
			"2019",
		},
	}

	t1 := time.Now()
	tasks := fs.NewTasks(pts, e)
	tasks1 := fs.NewTasks(pts1, e1)
	tasks2 := fs.NewTasks(pts2, e2)

	go func() {
		for _, f := range *tasks {
			_f := f
			go _f.Walk(c, done)
		}

		for _, f1 := range *tasks1 {
			_f1 := f1
			go _f1.Walk(c1, done)
		}

		for _, f2 := range *tasks2 {
			_f2 := f2
			go _f2.Walk(c1, done)
		}

	}()

	for {
		select {
		case v := <-c:
			num++
		case v := <-c1:
			num++
			go func() {
				xls <- fs.GetExpectedPath(v,`^[UCBP]2.*\.xlsx?`)
			}()
		case data := <-xls:
				//database operation
		case <-done:
			taskHaveDone++
			if taskHaveDone == total_len{
				fmt.Printf("updated %d 个 case\n", num)
				fmt.Printf("It takes %v in total\n", time.Since(t1))
				return
			}
		}
	}
}

```
