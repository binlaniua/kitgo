package db

import (
	"database/sql"
	"fmt"
	"github.com/binlaniua/kitgo"
	"github.com/binlaniua/kitgo/file"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)





//-------------------------------------
//
//
//
//-------------------------------------
func ConnectConfig(info *DataBaseConfig) *sql.DB  {
	if info.Alias == "" {
		info.Alias = DEFAULT_DB_NAME
	}
	return ConnectAsAlias(info.Alias, info.Ip, info.Port, info.User, info.Password, info.DB)
}

//-------------------------------------
//
//
//
//-------------------------------------
func ConnectFile(filePath string) *sql.DB {
	info := &DataBaseConfig{}
	file.LoadJsonFile(filePath, info)
	return ConnectConfig(info)
}

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
	loc, _ := time.LoadLocation("Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s", user, password, host, port, database, loc.String())
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		kitgo.ErrorLog.Panic("打开连接错误 => ", dsn, err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	err = db.Ping()
	if err != nil {
		kitgo.ErrorLog.Panic("打开连接错误 => ", dsn, err)
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
func DML(sql string, args ...interface{}) (sql.Result, error) {
	return DMLByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func DMLByAlias(alias string, sql string, args ...interface{}) (sql.Result, error) {
	db := GetDBByAlias(alias)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}
