package date

import (
	"time"
	"fmt"
)


//-------------------------------------
//
// 
//
//-------------------------------------
func NowDateStr() string {
	return NowDateStrExt("-")
}

func NowDateStrExt(join string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("2006%s01%s02", join, join))
}

//-------------------------------------
//
//
//
//-------------------------------------
func NowTimeStr() string {
	n := time.Now()
	return n.Format("15:04:05")
}

