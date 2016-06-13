package kitgo

import (
	"strings"
	"regexp"
	"errors"
)

//-------------------------------------
//
// 首字母小写
//
//-------------------------------------
func StringFirstLetterLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

//-------------------------------------
//
// 首字母大写
//
//-------------------------------------
func StringFirstLetterUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

//-------------------------------------
//
// 取首尾空格
//
//-------------------------------------
func StringTrim(src string) string {
	return strings.Trim(src, " \n")
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringBetween(src string, start string, end string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		src = src[sI + len(start):]
		eI := strings.Index(src, end)
		if eI > 0 {
			return src[:eI], true
		} else {
			Log(src, "开始位置[", end, "]没有找到")
			return src, false
		}
	} else {
		Log(src, "开始位置[", start, "]没有找到")
		return "", false
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringStartWith(src string, s string) bool {
	return strings.Index(src, s) == 0;
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringAfter(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		return src[sI + len(start):], true
	} else {
		return "", false
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringBefore(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 1 {
		return src[:sI], true
	} else {
		return "", false
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringMatch(src string, p string, group int) (string, error) {
	pattern, err := regexp.Compile(p)
	if err != nil {
		return "", err
	} else {
		r := pattern.FindStringSubmatch(src)
		if len(r) > group {
			return r[group], nil
		} else {
			return "", errors.New("匹配失败")
		}
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringLeftPad(src string, length int, pad string) string {
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
