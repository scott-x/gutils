/*
* @Author: scottxiong
* @Date:   2020-04-07 15:46:54
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-08-10 15:35:44
 */
package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/scott-x/gutils/model"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ParseErr = errors.New("parse config file error")
	DbError  = errors.New("default database can't be null")
)

func GetConnection(file string) (*sql.DB, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &model.DBConfig{}
	err = json.Unmarshal(bs, config)
	if err != nil {
		return nil, ParseErr
	}

	driver, username, password, host, port, database := config.Driver, config.Username, config.Password, config.Host, config.Port, config.Database

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
		return nil, DbError
	}
	dataSource := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8"
	return sql.Open(driver, dataSource)
}
