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
- `func ReadAndReplace(file string, replace map[string]string) error`: read file and replace it's content with map[string]string{"old1":"new1","old2":"new2"}
- `func ReadFile1(file string) (string, error)`: read the file and get the string content
- `func GetRunningFolder() string `: get running folder, same as linux's command `pwd`
- `func GetEnv(env string) string`:get the environment of the system
- `func NewFile(filePath string) `: create file