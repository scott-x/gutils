/*
* @Author: scottxiong
* @Date:   2020-04-14 20:57:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-08-10 15:52:02
 */
package parse

import (
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/model"
	"regexp"
	"strings"
)

var filed_keywords = []string{
	"INT",
	"VARCHAR",
	"DOUBLE",
	"ENUM",
	"DATETIME",
	"TEXT",
}

func GetTables(fileName string) *model.Tables {
	tables := &model.Tables{}
	str_arr := get_table_declaration(fileName)
	for _, v := range str_arr {

		table := &model.Table{}

		//fields
		arr := strings.Split(v, "\n")
		for _, v := range arr {
			if v[len(v)-1] == '(' {
				//table start
				table.Name = getTableName(v)

			} else if strings.Contains(v, ";") {
				//table end

			} else {
				x, y := getType(v)

				(*table).Fields = append((*table).Fields, model.Field{x, y})
			}
		}
		*tables = append(*tables, *table)
	}
	return tables
}

func get_table_declaration(fileName string) []string {
	c, _ := fs.ReadFile1(fileName)
	//remove comment start with # or --
	re_comment1 := regexp.MustCompile(`(\s{0,}--.*)`)
	re_comment2 := regexp.MustCompile(`(\s{0,}#.*)`)
	// re_comment3 := regexp.MustCompile(`(\s{0,}comment.*),`)
	// re_comment4 := regexp.MustCompile(`(\s{0,}comment.*)`)
	c = re_comment1.ReplaceAllString(c, "")
	c = re_comment2.ReplaceAllString(c, "")
	// c = re_comment3.ReplaceAllString(c, ",")
	// c = re_comment4.ReplaceAllString(c, "")

	//replace create table => CREATE TABLE
	re_replace := regexp.MustCompile(`(create table)`)
	c = re_replace.ReplaceAllString(c, "CREATE TABLE")
	// fmt.Println(c)
	// get the table comment
	re := regexp.MustCompile(`CREATE TABLE(.*\n)(.*,\n)+(.*\n).*;`)
	return re.FindAllString(c, -1)
}

func getType(str string) (string, string) {
	//id VARCHAR(64) NOT NULL PRIMARY KEY,
	var res string
	arr := strings.Split(strings.TrimSpace(str), " ")
	fileName := arr[0]
	filedType := arr[1]
	filedType = strings.ToUpper(filedType)
	if strings.Contains(filedType, "INT") || strings.Contains(filedType, "ENUM") { //enum => 0,1
		res = "int"
	} else if strings.Contains(filedType, "VARCHAR") || strings.Contains(filedType, "CHAR") || strings.Contains(filedType, "TEXT") {
		res = "string"
	} else {
		res = "string"
	}
	return fileName, res
}

func getTableName(str string) string {
	//CREATE TABLE comment(
	str = strings.Split(str, "(")[0]
	arr := strings.Split(str, " ")
	return arr[len(arr)-1]
}
