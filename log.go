package kitgo

import (
	"log"
	"os"
)

//-------------------------------------
//
//
//
//-------------------------------------
var (
	DebugLog *log.Logger = log.New(os.Stdout, "", log.Ltime|log.Ldate|log.Llongfile)
	ErrorLog *log.Logger  = log.New(os.Stderr, "", log.Ltime|log.Ldate|log.Llongfile)
)

//
//
//
//
//
func init() {
}
