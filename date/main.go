package date

import (
	"fmt"
	"strconv"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
func DateToString() string {
	return DateToStringByJoin("-")
}

//-------------------------------------
//
//
//
//-------------------------------------
func DateToStringByJoin(join string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("2006%s01%s02", join, join))
}

//-------------------------------------
//
//
//
//-------------------------------------
func DateTimeToStringNoJoin() string {
	n := time.Now()
	return n.Format("20060102150405")
}

//-------------------------------------
//
//
//
//-------------------------------------
func DateTimeToString() string {
	return DateTimeToStringByJoin("-", ":")
}

//-------------------------------------
//
//
//
//-------------------------------------
func DateTimeToStringByJoin(dateJoin string, timeJoin string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("2006%s01%s02 15%s04%s05", dateJoin, dateJoin, timeJoin, timeJoin))
}

//-------------------------------------
//
//
//
//-------------------------------------
func TimeToString() string {
	n := time.Now()
	return n.Format("15:04:05")
}

//-------------------------------------
//
//
//
//-------------------------------------
func TimeToStringByJoin(join string) string {
	n := time.Now()
	return n.Format(fmt.Sprintf("15%s04%s05", join, join))
}

//-------------------------------------
//
//
//
//-------------------------------------
func TimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
