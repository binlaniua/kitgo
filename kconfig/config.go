package kconfig

import (
	"github.com/binlaniua/kitgo/kfile"
	"encoding/json"
	"log"
)

var (
	Debug = false
)
//-------------------------------------
//
//
//
//-------------------------------------
func Log(args ... interface{}) {
	if Debug {
		log.Println(args)
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func LogF(f string, args ... interface{}) {
	if Debug {
		log.Panicf(f, args...)
	}
}

//-------------------------------------
//
//  读取Json文件
//
//-------------------------------------
func LoadJsonFile(filePath string) bool {
	data, ok := kfile.ReadBytes(filePath)
	if !ok {
		return false
	}
	err := json.Unmarshal(data, c)
	if err != nil {
		Log(filePath, "加载失败 => ", err)
		return false
	}
	return true
}

//-------------------------------------
//
//  读取配置文件
//
//-------------------------------------
func LoadProperties(filePath string) bool {
	return
}
