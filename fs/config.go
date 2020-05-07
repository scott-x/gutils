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

func ModifyAttrOfJson(filename string, key string, value interface{}) {
	v := ReadJson(filename)(key)
	if v == nil {
		log.Printf("%s doesn't exist", key)
		return
	}
	old := ReadJson(filename)(key).(string)
	m := make(map[string]interface{}, 0)
	// m[old] = fmt.Sprintf("%f", value)
	m[old] = value
	err := ReadAndReplace(filename, m)
	if err != nil {
		panic(err)
	}
}
