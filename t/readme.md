# package t

### API
- `func GetRecentDays(Days int) []string `: Days can be negative value
- `func GetTime(type_t string) string`: get time

### options of type_t

```
type_t		result format
-------------------------------------------------------
"yyyy-mm-dd hh:mm:ss"	2006-01-02 15:04:05
"yyyymmdd"           	"20060102"
"yyyy年mm月dd日"      	"2006年01月02日"
"yyyy-mm-dd"         	"2006-01-02"
"yyyy/mm/dd"         	"2006/01/02"
"mmdd"               	"0102"
"mm/dd"              	"01/02"
"mm-dd"              	"01-02"
"mm月dd日"            	"01月02日"
"hh:mm:ss"           	"15:04:05"
"hh:mm"              	15:04"
"hh时mm分"            	"15时04分"
"hh时mm分ss秒"        	"15时04分05秒"
""                   	"2006-01-02 15:04:05"
"yyyy":              	"2006"	
"yyyy年"             	"2006年"
"yyyymm"		"200601"
```
