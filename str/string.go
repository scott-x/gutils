/*
* @Author: scottxiong
* @Date:   2019-09-09 16:59:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-09 17:56:48
 */
package str

import (
	"fmt"
	"github.com/scott-x/gutils/model"
	"strings"
)

func FirstLetterToUpper(str string, mod int) string {
	w := ""
	switch mod {
	case 1:
		w = strings.ToUpper(str[0:1]) + str[1:]
	case -1:
		arr := strings.Split(str, " ")
		for _, v := range arr {
			if strings.ToLower(v)[0] >= 97 && strings.ToLower(v)[0] <= 122 {
				w += strings.ToUpper(v[:1]) + v[1:] + " "
			} else {
				w += v + " "
			}

		}
	}
	return strings.Trim(w, " ")
}

func FindAllSubPositions(str string, sub string) []int {
	points := make([]int, 0)
	i := strings.Index(str, sub)
	for i != -1 {
		points = append(points, i)
		str = str[0:i] + strings.Repeat("*", len(sub)) + str[i+len(sub):]
		i = strings.Index(str, sub)
	}
	return points
}

func GetWord(str string, i int) string {
	arr := strings.Split(str, " ")
	if i <= len(arr) {
		return arr[i+1]
	} else {
		return ""
	}

}

//辅助函数
func RangeRune(str string) {
	for k, v := range []rune(str) {
		fmt.Printf("%d->%v ", k, v)
	}
	fmt.Printf("\n")
}

//找出所有sub string出现的位置(最短的): start with first, end with last
func GetPositions(content string, first, last string) *model.Ps {
	ps := &model.Ps{}
	top := FindAllSubPositions(content, first)
	end := FindAllSubPositions(content, last)
	for _, x := range top {
		for _, y := range end {
			if y > x {
				p := &model.P{}
				p.Start = x
				p.End = y + 1
				*ps = append(*ps, *p)
				break
			}
		}
	}
	return ps
}
