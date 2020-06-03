/*
* @Author: scottxiong
* @Date:   2020-06-03 11:47:16
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-06-03 11:55:23
 */
package fs

import (
	"github.com/scott-x/gutils/str"
)

/*
we can use md5 to check whether file was updated or not?
we can also used for file checksum.
*/
func MD5(file string) string {
	content, err := ReadFile1(file)
	if err != nil {
		panic(err)
	}
	return str.MD5(content)
}
