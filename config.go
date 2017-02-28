package kitgo

import (
	"log"
	"os"
)

var (
	DebugLog *log.Logger
	ErrorLog *log.Logger
)

func init() {
	DebugLog = log.New(os.Stdout, "", log.Ltime|log.Ldate|log.Llongfile)
	ErrorLog = log.New(os.Stderr, "", log.Ltime|log.Ldate|log.Llongfile)
}

//-------------------------------------
//
//
//
//-------------------------------------
func Log(args ...interface{}) {
}

//-------------------------------------
//
//
//
//-------------------------------------
func LogF(f string, args ...interface{}) {
}
