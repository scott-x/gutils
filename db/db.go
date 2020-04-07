/*
* @Author: scottxiong
* @Date:   2020-04-07 15:46:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-07 16:29:31
 */
package db

import (
	"database/sql"
	"github.com/scott-x/gutils/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(config *model.DBConfig) (*sql.DB, error) {
	driver := config.Driver
	username := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	database := config.Database

	if driver == "" {
		driver = "mysql"
	}

	if username == "" {
		username = "root"
	}

	if password == "" {
		password = "root"
	}

	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "3306"
	}

	if database == "" {
		database = "test"
	}

	return sql.Open(driver, username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8")
}
