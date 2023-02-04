/*
* @Author: scottxiong
* @Date:   2019-07-25 16:15:00
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 19:12:03
 */
package fs

import (
	"testing"
)

func TestList(t *testing.T) {
	//just list files & folders current folder
	res := List("example")
	var expected = []string{"js","1.txt","2.txt"}
	if len(res)!=len(expected) {
		t.Errorf("Test List(\"example\") failed, expected [\"js\",\"1.txt\",\"2.txt\"], but got %v",res)
	}
}

func TestReadAndReplace(t *testing.T) {
	//just list files & folders current folder
	m:=make(map[string]string)
	m["Test"] = "Movie"
	m["test"] = "movie"
	Copy("example/js/test.js","example/js/movie.js")
	ReadAndReplace("example/js/movie.js",m)
}
