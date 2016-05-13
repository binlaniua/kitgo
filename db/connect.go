package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

//
// 所有应用共享实例
//
var DBMap map[string]*sql.DB = make(map[string]*sql.DB)
var DEFAULT_DB_NAME = "___default"

//-------------------------------------
//
// 默认连接
//
//-------------------------------------
func Connect(host, port, user, password, db string) *sql.DB {
	return ConnectAsAlias(DEFAULT_DB_NAME, host, port, user, password, db)
}

//-------------------------------------
//
// 多库连接
//
//-------------------------------------
func ConnectAsAlias(alias, host, port, user, password, database string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("打开连接错误 => ", dsn, err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	err = db.Ping()
	if err != nil {
		log.Println("打开连接错误 => ", dsn, err)
	}
	DBMap[alias] = db
	return db
}

//-------------------------------------
//
// 获取默认DB
//
//-------------------------------------
func GetDB() *sql.DB {
	r := GetDBByAlias(DEFAULT_DB_NAME)
	return r
}

//-------------------------------------
//
// 多库连接
//
//-------------------------------------
func GetDBByAlias(alias string) *sql.DB {
	db, ok := DBMap[alias]
	if ok {
		return db
	} else {
		log.Panic("获取DB失败 =>", alias)
		return nil
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func DML(sql string, args ... interface{}) (sql.Result, error) {
	return DMLByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func DMLByAlias(alias string, sql string, args ... interface{}) (sql.Result, error) {
	db := GetDBByAlias(alias)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}