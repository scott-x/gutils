/*
* @Author: scottxiong
* @Date:   2020-04-14 20:57:49
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-15 00:01:39
 */
package parse

import (
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/model"
	"strings"
	"regexp"
)

var filed_keywords = []string{
	"INT",
	"VARCHAR",
	"DOUBLE",
	"ENUM",
	"DATETIME",
	"TEXT",
}

func GetTables(fileName string) *model.Tables{
	tables := &model.Tables{}
	str_arr := get_table_declaration(fileName)
	for _, v := range str_arr {

	    table := &model.Table{}

	    //fields
	    arr := strings.Split(v,"\n")
	    for _,v := range arr {
	    	if v[len(v)-1]=='(' {
	    		//table start
	    		table.Name = getTableName(v)
	             
	    	}else if strings.Contains(v,";") {
	    		//table end

	    	}else {
	    		x, y := getType(v)

	    		(*table).Fields = append((*table).Fields,model.Field{x,y})
	    	}
	    }
	    *tables  = append(*tables,*table)
	}
	return tables
}

func get_table_declaration(fileName string) []string{
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

// func GetTables(fileName string) *model.Tables {
// 	tables := &model.Tables{}
// 	var newArr []string
// 	content := fs.ReadFile(fileName)
// 	//按行存入数组中
// 	arr := strings.Split(content, "\n")
// 	for _, x := range arr {
// 		if strings.Contains(strings.ToUpper(x), "CREATE TABLE") && strings.Contains(x, "(") {
// 			newArr = append(newArr, x)
// 		}
// 		//过滤注释 # --
// 		if strings.HasPrefix(strings.TrimSpace(x), "--") || strings.HasPrefix(strings.TrimSpace(x), "#") || strings.Contains(strings.ToUpper(x), "INSERT") {

// 		} else if strings.Contains(x, ";") {
// 			if strings.Contains(strings.ToUpper(x), "DATABASE") || strings.Contains(strings.ToUpper(x), "UPDATE") || strings.Contains(strings.ToUpper(x), "DELETE") || strings.Contains(strings.ToUpper(x), "DROP") {

// 			}

// 		} else {

// 			for _, y := range filed_keywords {
// 				// 如果包含mysql type关键字 或者 有 CREATE TABLE comment( 等字样则保留
// 				if strings.Contains(strings.ToUpper(x), y) {
// 					newArr = append(newArr, x)
// 				}
// 			}

// 		}
// 	}
// 	table := &model.Table{}
// 	for _, z := range newArr {

// 		//table start
// 		if strings.Contains(strings.ToUpper(z), "CREATE TABLE") {
// 			//add table
// 			if table.Name != "" {
// 				*tables = append(*tables, *table)
// 			}
// 			//reset table to nil
// 			table = &model.Table{}
// 			tb := getTableName(z)
// 			table.Name = tb
// 			// fmt.Println(tb)
// 		} else {
// 			f := &model.Field{}

// 			field, t := getType(z)
// 			f.Name = field
// 			f.Type = t
// 			// fmt.Println("field:" + field + ", type:" + t)
// 			(*table).Fields = append((*table).Fields, *f)
// 		}
// 		//table end
// 	}
// 	return tables
// }

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
