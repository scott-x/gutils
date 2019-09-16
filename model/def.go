/*
* @Author: scottxiong
* @Date:   2019-09-05 21:47:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-16 11:21:19
 */
package model

type Question struct {
	Name  string
	Tip   string
	ReTip string
	Re    string
	//Do func()
}

type SimpleQuestion struct {
	Tip string
}
type Questions struct {
	Qs []Question
}

type Tasks struct {
	Names []string
}

type P struct {
	Start int
	End   int
}

type Ps []P
