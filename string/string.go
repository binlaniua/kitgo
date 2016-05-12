package string

import (
	"strings"
	"regexp"
	"github.com/binlaniua/kitgo/config"
)

//-------------------------------------
//
// 首字母小写
//
//-------------------------------------
func FirstLetterLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

//-------------------------------------
//
// 首字母大写
//
//-------------------------------------
func FirstLetterUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

//-------------------------------------
//
// 取首尾空格
//
//-------------------------------------
func Trim(src string) string {
	return strings.Trim(src, " ")
}

//-------------------------------------
//
// 
//
//-------------------------------------
func Between(src string, start string, end string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		src = src[sI + len(start):]
		eI := strings.Index(src, end)
		if eI > 0 {
			return src[:eI], true
		} else {
			config.Log(src, "开始位置[", end, "]没有找到")
			return src, false
		}
	} else {
		config.Log(src, "开始位置[", start, "]没有找到")
		return "", false
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StartWith(src string, s string) bool {
	return strings.Index(src, s) == 0;
}

//-------------------------------------
//
// 
//
//-------------------------------------
func After(src string, start string) string {
	sI := strings.Index(src, start)
	if sI >= 0 {
		return src[sI + len(start):]
	} else {
		return "";
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func Before(src string, start string) string {
	sI := strings.Index(src, start)
	if sI >= 1 {
		return src[:sI]
	} else {
		return "";
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func Match(src string, p string, group int) string {
	pattern, err := regexp.Compile(p)
	if err != nil {
		config.Log("构建正则出错 =>", p)
		return ""
	} else {
		r := pattern.FindStringSubmatch(src)
		if len(r) > group {
			return r[group]
		} else {
			return ""
		}
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func LeftPad(src string, length int, pad string) string {
	srcLen := len(src)
	offset := length - srcLen
	for i := 0; i < offset; i++ {
		src = pad + src
	}
	return src
}

//-------------------------------------
//
// 
//
//-------------------------------------
