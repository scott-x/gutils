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
- `func ReadAndReplace(file string, replace map[string]interface{}) error`: read file and replace it's content with `map[string]string{"old1":"new1","old2":"new2"}`
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
- `func ListAll1(folders []string, ignore []string) (*FS, int64, error)`:
- `func Zip(zipName string, Base string, files []string) `: zip a file, `Base` will be removed.
- `func ZipWithBar(z *ZIP) `: zip with progress bar
- `func Tab(n int) string`: tabale n
- `func ReadJson(filename string) func(string) interface{}`: read the configuration of json, return a method, with which we can get the related value.

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
	Ignore  []string
	Where   string
	Base    string // will be truncated from the full path
}
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
		"/Volumes/datavolumn_bmkserver_Pub/202005/0506/C2002H1_PMG/AI_ThisFolderToPrinter",
		"/Volumes/datavolumn_bmkserver_Pub/202005/0506/C2002H1_PMG/PDF_Locked_For_Visual_Ref",
	}

	ignore := []string{".git"}

	z := &fs.ZIP{
		folders,
		ignore,
		"/Users/apple/Desktop/C2002H1_PMG.zip",
		"/Volumes/datavolumn_bmkserver_Pub/202005/0506/C2002H1_PMG",
	}

	fs.ZipWithBar(z)
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
  }
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
	urlValue := get("mysql.url")
	fmt.Println("mysql url:", urlValue)
	fmt.Printf("mysql url: %s\nmysql username: %s\nmysql password: %s\n", get("mysql.url"), get("mysql.username"), get("mysql.password"))
}
```