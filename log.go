package kitgo

import (
	"log"
	"os"
)

//-------------------------------------
//
// 文件日志
//
//-------------------------------------
func NewFileLog(filePath string) *log.Logger {
	debugFile, err := os.OpenFile(filePath, os.O_CREATE | os.O_APPEND, 0x766);
	if err != nil {
		log.Panicf("创建日志文件失败 => [%v]", err)
	}
	return log.New(debugFile, "", log.Llongfile | log.Ltime | log.Ldate)
}