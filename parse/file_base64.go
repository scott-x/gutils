/*
* @Author: scottxiong
* @Date:   2020-04-28 17:17:25
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-28 17:37:20
 */
package parse

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

//file => base64
func GetBase64FromFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(data)
	return sourcestring
}
