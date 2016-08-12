package date

import (
	"time"
	"fmt"
	"strconv"
)

//-------------------------------------
//
//
//
//-------------------------------------
func DateTimeNoJoinString() string {
	n := time.Now()
	return n.Format("20060102150405")
}

//-------------------------------------
//
// 
//
//-------------------------------------
func NowDateStr() string {
	return NowDateStrExt("-")
}

//-------------------------------------
//
//
//
//-------------------------------------
func NowDateStrExt(join string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("2006%s01%s02", join, join))
}

//-------------------------------------
//
// 获取当前小时
//
//-------------------------------------
func GetHour() int {
	n := time.Now()
	r := n.Format("15")
	h, _ := strconv.Atoi(r)
	return h
}

//-------------------------------------
//
// 获取当前分钟
//
//-------------------------------------
func GetMinute() int {
	n := time.Now()
	r := n.Format("04")
	h, _ := strconv.Atoi(r)
	return h
}

//-------------------------------------
//
//
//
//-------------------------------------
func NowDateTimeStr() string {
	return NowDateTimeStrExt("-", ":")
}

func NowDateTimeStrExt(dateJoin string, timeJoin string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("2006%s01%s02 15%s04%s05", dateJoin, dateJoin, timeJoin, timeJoin))
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

//-------------------------------------
//
// 
//
//-------------------------------------
func TimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
