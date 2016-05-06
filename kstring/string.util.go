package kstring

import (
	"strings"
	"log"
	"regexp"
	"net/url"
)

//-------------------------------------
//
//
//
//-------------------------------------
func StringFirstLetterLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringFirstLetterUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringTrim(src string) string {
	return strings.Trim(src, " ")
}

//-------------------------------------
//
// 
//
//-------------------------------------
func StringBetween(src string, start string, end string) string {
	sI := strings.Index(src, start)
	if sI >= 0 {
		src = src[sI + len(start):]
		eI := strings.Index(src, end)
		if eI > 0 {
			return src[:eI]
		} else {
			return src
		}
	} else {
		return "";
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
func StringAfter(src string, start string) string {
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
func StringBefore(src string, start string) string {
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
func StringMatch(src string, p string, group int) string {
	pattern, err := regexp.Compile(p)
	if err != nil {
		log.Println("构建正则出错 =>", p)
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
func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}