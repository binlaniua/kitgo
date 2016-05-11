package kconfig

import (
	"log"
)

var (
	Debug = true
)

func init() {
	log.SetFlags(log.Lshortfile)
}

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