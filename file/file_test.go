package file

import (
	"log"
	"testing"
	"time"
)

func Test_write_rotate(t *testing.T) {
	fr, err := NewFileRotate(&FileRotateOption{
		docPath: "c:\\",
		Prefix:  "si_",
		Suffix:  "_error.log",
		//rotate: "@every 1m",
		//format: "2006_01_02_15_04",
	})
	if err != nil {
		log.Print("111")
	}
	logger := log.New(fr, "", log.Llongfile)
	for i := 0; i < 200; i++ {
		logger.Println("1111")
		time.Sleep(time.Second)
	}
}
