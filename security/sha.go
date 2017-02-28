package security

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
func SHA1(s string) []byte {
	t := sha1.New()
	io.WriteString(t, s)
	return t.Sum(nil)
}

//-------------------------------------
//
//
//
//-------------------------------------
func SHA1Hex(s string) string {
	return hex.EncodeToString(SHA1(s))
}

//-------------------------------------
//
//
//
//-------------------------------------
func SHA1Map(m map[string]string, joinStr string, other ...string) string {
	//1 排序
	kList := make([]string, 0, len(m))
	for k, _ := range m {
		kList = append(kList, k)
	}
	sort.Strings(kList)

	//2 按 = 号相加
	strList := make([]string, 0, len(m))
	for _, k := range kList {
		v := fmt.Sprintf("%s=%s", k, m[k])
		strList = append(strList, v)
	}

	//3 加额外的
	allStr := strings.Join(strList, joinStr)
	if other != nil {
		for _, otherItem := range other {
			allStr = allStr + otherItem
		}
	}
	//4. 大写加密
	r := strings.ToUpper(SHA1Hex(allStr))
	return r
}
