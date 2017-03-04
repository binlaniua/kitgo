package db

import (
	"database/sql"
	"log"
	"github.com/binlaniua/kitgo"
)

//
// 所有应用共享实例
//
var (
	DBMap       map[string]*sql.DB = make(map[string]*sql.DB)
	debugLogger                    = kitgo.DebugLog
	errorLogger                    = kitgo.ErrorLog
)

const (
	DEFAULT_DB_NAME = "___default"
)

//-------------------------------------
//
//
//
//-------------------------------------
func SetDebugLogger(l *log.Logger) {
	debugLogger = l
}

//-------------------------------------
//
//
//
//-------------------------------------
func SetErrorLogger(l *log.Logger) {
	errorLogger = l
}

//-------------------------------------
//
//
//
//-------------------------------------
type DataBaseConfig struct {
	Alias    string `json:"alias"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}
