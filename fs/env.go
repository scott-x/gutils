/*
* @Author: scottxiong
* @Date:   2019-12-27 01:40:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-07 01:08:03
 */
package fs

import (
	"os"
)

var (
	DESKTOP = GetEnv("HOME") + "/Desktop"
	HOME    = GetEnv("HOME")
)

func GetEnv(env string) string {
	e, flag := os.LookupEnv(env)
	if flag {
		return e
	} else {
		return ""
	}
}
