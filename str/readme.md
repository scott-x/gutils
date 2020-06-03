### Attention
When dealing with string, we'd better convert it to ASCII 1st，for some characters can't be seen，such as `\n`

### Common ASCII

```
space			32
enter			10
tab			9
*			42
0			48
A			65
a 			97
```
for more information, please [click here](http://www.asciitable.com/)

### API
- `func FirstLetterToUpper(str string, mod int) string`: mod can be set 1(for single world) or -1(for all words)
- `func FindAllSubPositions(str string, sub string) []int`:
- `func GetWord(str string, i int) string `: i starts from 1
- `func RangeRune(str string)`: the function is used for development.
- `func GetPositions(content string, first, last string) *model.Ps`;
- `func IsLastItem(arr []string, index int) bool`: judge if it's the last item of the Array
- `func GetContentBetween(reource string, A string, B string) string `: get the subcontent between A and B, if not match, return `""` 
- `func MD5(str string) string`: get the md5 string(with fixed length), note that md5 algorithem is irreversible.

### rune
```
#golang中string取索引值得到的是unicode，unicode兼容ASCII，比ASCII范围更广
a := "hello"
fmt.Println(a[0]) //104

s := "你好啊"
fmt.Println(s[0]) //228
```

### `func GetContentBetween(reource string, A string, B string) string`
```golang
package main

import (
	"fmt"
	"github.com/scott-x/gutils/str"
)

func main() {
	sub := str.GetContentBetween("hello (wold(a,b) yds", "wold(", ")")
	fmt.Println(sub)
}
```
```
a,b
```