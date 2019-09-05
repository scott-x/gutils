/*
* @Author: scottxiong
* @Date:   2019-09-05 21:47:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-09-05 21:58:14
 */
package model

type Question struct {
	Name  string
	Tip   string
	ReTip string
	Re    string
	//Do func()
}

type Questions struct {
	Qs []Question
}

type Tasks struct {
	Names []string
}
