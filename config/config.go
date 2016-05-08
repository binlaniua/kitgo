package config

import "log"

var (
	Debug = false
)

func Log(args ... interface{}) {
	log.Println(args)
}

func LogF(f string, args ... interface{}) {
	log.Panicf(f, args...)
}