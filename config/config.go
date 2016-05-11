package config

import "log"

var (
	Debug = false
)

func Log(args ... interface{}) {
	if Debug {
		log.Println(args)
	}
}

func LogF(f string, args ... interface{}) {
	if Debug {
		log.Panicf(f, args...)
	}
}