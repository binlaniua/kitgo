package kitgo

import (
	"strings"
	"regexp"
	"errors"
)

//-------------------------------------
//
// 正则切割
//
//-------------------------------------
func StringSplitByRegexp(src string, reg string) []string  {
	pattern := regexp.MustCompile(reg)
	indexes := pattern.FindAllStringIndex(src, -1)
	laststart := 0
	result := make([]string, len(indexes) + 1)
	for i, element := range indexes {
		result[i] = src[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = src[laststart:len(src)]
	return result
}

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
	return strings.TrimSpace(src)
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
