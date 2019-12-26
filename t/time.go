/*
* @Author: scottxiong
* @Date:   2019-12-27 01:52:12
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 02:28:00
 */
package t

import (
	"time"
)

func GetRecentDays(Days int) []string {
	currentTime := time.Now()
	var times []time.Time
	var arr []string
	for i := 0; i < Days; i++ {
		times = append(times, currentTime.AddDate(0, 0, i))
	}

	for _, t := range times {
		arr = append(arr, t.Format("0102"))
	}
	return arr
}

func GetTime(type_t string) string {
	var t string
	time1 := time.Now()
	switch type_t {
	case "yyyy-mm-dd hh:mm:ss":
		t = time1.Format("2006-01-02 15:04:05")
	case "yyyymmdd":
		t = time1.Format("20060102")
	case "yyyy年mm月dd日":
		t = time1.Format("2006年01月02日")
	case "yyyy-mm-dd":
		t = time1.Format("2006-01-02")
	case "yyyy/mm/dd":
		t = time1.Format("2006/01/02")
	case "mmdd":
		t = time1.Format("0102")
	case "mm/dd":
		t = time1.Format("01/02")
	case "mm-dd":
		t = time1.Format("01-02")
	case "mm月dd日":
		t = time1.Format("01月02日")
	case "hh:mm:ss":
		t = time1.Format("15:04:05")
	case "hh:mm":
		t = time1.Format("15:04")
	case "hh时mm分":
		t = time1.Format("15时04分")
	case "hh时mm分ss秒":
		t = time1.Format("15时04分05秒")
	default:
		t = time1.Format("2006-01-02 15:04:05")
	}
	return t
}
