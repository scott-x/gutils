/*
* @Author: scottxiong
* @Date:   2020-05-08 03:16:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-08 07:11:52
 */
package fs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

/*
   1. read and modify the json data
*/

//read the configuration of json, return a method, with which we can get the related value
func ReadJson(filename string) func(string) interface{} {
	dir := path.Dir(filename)
	base := path.Base(filename)

	//set the name of the file(removed .json)
	viper.SetConfigName(base[:len(base)-5])
	//set config path
	viper.AddConfigPath(dir)
	//set file type is json
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	return viper.Get
}

//modify string value
func ModifyAttrOfJson_STRING(filename string, key string, value string) {
	v := ReadJson(filename)(key)
	if v == nil {
		log.Printf("%s doesn't exist", key)
		return
	}
	old := ReadJson(filename)(key).(string)
	m := make(map[string]string, 0)
	// m[old] = fmt.Sprintf("%f", value)
	m[old] = value
	err := ReadAndReplace(filename, m)
	if err != nil {
		panic(err)
	}
}

// modify int value
func ModifyAttrOfJson_FLOAT64(filename string, key string, value float64) {
	v := ReadJson(filename)(key)
	if v == nil {
		log.Printf("%s doesn't exist", key)
		return
	}
	old := ReadJson(filename)(key).(float64)
	m := make(map[string]string, 0)
	// m[old] = fmt.Sprintf("%f", value)
	m[float64_to_string(old)] = float64_to_string(value)
	err := ReadAndReplace(filename, m)
	if err != nil {
		panic(err)
	}
}
