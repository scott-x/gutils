/*
* @Author: scottxiong
* @Date:   2020-05-08 03:16:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-05-08 03:31:19
 */
package fs

import (
	"fmt"
	"github.com/spf13/viper"
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
