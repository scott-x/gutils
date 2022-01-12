/*
* @Author: apple
* @Date:   2022-01-12 20:55:50
* @Last Modified by:   scottxiong
* @Last Modified time: 2022-01-12 21:48:32
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/scott-x/gutils/cmd"
)

type Wifi struct {
	Ip           string `json:"ip"`
	Delay        string `json:"delay"`
	DowloadSpeed string `json:"dowload_speed"`
}

type Wifis []Wifi

func (wifis Wifis) HandleItems() []string {
	var items []string
	for _, wifi := range wifis {
		items = append(items, fmt.Sprintf("\t%s\t%s\t\t%s", wifi.Ip, wifi.Delay, wifi.DowloadSpeed))
	}
	return items
}

func getWifis() []Wifi {
	var wifis []Wifi
	bs, _ := ioutil.ReadFile("./wifi.json")
	_ = json.Unmarshal(bs, &wifis)
	return wifis
}

func main() {
	wifis := getWifis()
	num := cmd.SelectOne("Good ip as following", fmt.Sprintf("NO\tIP\t\tDelay(ms)\tSpeed(MB/s)"), 7, Wifis(wifis))
	fmt.Println(num)
}
