/*
* @Author: scottxiong
* @Date:   2019-07-25 18:45:44
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 18:49:00
 */
//https://learngolang.net/tutorials/how-to-remove-all-files-in-a-directory-in-go/
package fs

import (
	"fmt"
	"os"
)

func RemoveAll(folder string) {
	err := os.RemoveAll(folder)
	fmt.Println(err)
}
