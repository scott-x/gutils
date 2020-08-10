/*
* @Author: scottxiong
* @Date:   2019-09-05 21:47:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-08-10 16:02:43
 */
package model

// the postion that the copy will be inserted
const (
	POSITION_BEFORE       = iota //above
	POSITION_AFTER               //below
	POSITION_CURRENT_LINE        //current line will be replaced
)

// aliyun oss
type OSS struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

//database configuration
type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

//question
type Question struct {
	Name  string
	Tip   string
	ReTip string
	Re    string
	//Do func()
}

//bref version
type SimpleQuestion struct {
	Tip string
}

//questions
type Questions struct {
	Qs []Question
}

//task
type Tasks struct {
	Names []string
}

//postion
type P struct {
	Start int
	End   int
}

//postions
type Ps []P

//insert structure
type Insert struct {
	File     string
	Content  string //new included content
	Postion  int    //before or after
	Line     int    //which line
	Keywords string //locate the line as per keywords that be given
	Replace         // optional, if you want to exec replace operation, add it on
}

//replace
type Replace struct {
	Old string
	New string
}
