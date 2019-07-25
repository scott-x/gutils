# package fs

### API
- `func Dir() string`: return the exec binary file folder
- `func RemoveAll(folder string)`: rm -rf
- `func Rename(oldName, newName string)`:rename the file or folder
- `func WriteBuf(file, data string) `: wirte a string data into the file
- `func WriteString(file, data string) `wirte a string data into the file
- `func Copy(inputFile, outputFile string) `: copy file
- `func CopyFile(srcName, dstName string) (written int64, err error)`:copy file
- `func CopyAndReplace(inputFile, outputFile string, replace map[string]string)`:copy and replace
- `func CreateDirIfNotExist(dir string)`: similar like `mkdir -p /a/b/c`
- `func CopyFolder(src, des string) error`:copy folder, for example: `CopyFolder("/Users/apple/go/src/github.com/scott-x/gutils", "/Users/apple/desktop/a")` :copy all files from gutils to folder `a`, not include folder `gutils`

