/*
* @Author: scottxiong
* @Date:   2019-12-27 01:40:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 01:43:31
 */
package fs

import (
	"os"
)

func GetEnv(env string) string {
	e, flag := os.LookupEnv(env)
	if flag {
		return e
	} else {
		return ""
	}
}
