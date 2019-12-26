### Attention
When dealing with string, we'd better convert to ASCII，for some characters con't be seen，such as `\n`

### Common ASCII
- space 32
- enter 10

### API
- `func FirstLetterToUpper(str string, mod int) string`: mod can be set 0 or 1
- `func FindAllSubPositions(str string, sub string) []int`:
- `func GetWord(str string, i int) string `: i starts from 1
- `func RangeRune(str string)`: the function is used for development.
- `func GetPositions(content string, first, last string) *model.Ps`;

### rune
```
a := "hello"
fmt.Println(a[0]) //104
```