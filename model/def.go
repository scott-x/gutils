/*
* @Author: scottxiong
* @Date:   2019-09-05 21:47:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-07 16:54:37
 */
package model

type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	PoolSize int
}

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

type Insert struct {
	File     string
	NewLine  string
	Line     int
	Keywords string
	Replace
}

type Replace struct {
	Old string
	New string
}
