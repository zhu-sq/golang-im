package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	rootDbPwd = "1234"
)

func init() {
	var err error

	connStr := "root:" + rootDbPwd + "@/mysql?charset=utf8&loc=Local&parseTime=true"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("sql open failed" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("sql ping failed" + err.Error())
	}
	crDb := "CREATE DATABASE IF NOT EXISTS golangim DEFAULT CHARSET utf8 COLLATE utf8_general_ci"
	stmt, err := db.Prepare(crDb)
	if err != nil {
		log.Fatal("sqlPrepare failed" + err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	connStr = "root:" + rootDbPwd + "@/golangim?charset=utf8&loc=Local&parseTime=true"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("sqlopen failed" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("sqlping failed" + err.Error())
	}

	//创建用户信息表
	cr_table := `create table if not exists user(
		 uID int(64) auto_increment primary key,
		 account varchar(25) unique,
		 pwd varchar(128) not null,
		 email varchar(20),
		 phone int(20),
		 nickname varchar(50),
		 icon varchar(256),
		 sex enum('男','女','保密') default '保密',
		 birthday date,
		 addr varchar(256)
		 )`
	stmt, err = db.Prepare(cr_table)
	if err != nil {
		log.Fatal("sqlPrepare failed" + err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}

}

func GetDB() *sql.DB {
	return db
}
