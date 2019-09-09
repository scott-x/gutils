### 注意
处理字符串，最好转化为ASCII处理，因为字符串有些东西是看不见的，容易忽略，比如`\n`

### 常见ASCII
- space 32
- enter 10

### API
- `func FirstLetterToUpper(str string, mod int) string`: mod 可以取0，1
- `func FindAllSubPositions(str string, sub string) []int`:
- `func GetWord(str string, i int) string `: i从1开始
- `func RangeRune(str string)`: 辅助函数，打印调试用
- `func GetPositions(content string, first, last string) *model.Ps`;

### rune
```
a := "hello"
fmt.Println(a[0]) //104
```