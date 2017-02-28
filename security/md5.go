package security

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//-------------------------------------
//
//
//
//-------------------------------------
func MD5WithSalt(s string, salt string) string {
	salt = MD5(salt)
	return MD5(s + salt)
}

//-------------------------------------
//
// 有序MD5加密
//
//-------------------------------------
func MD5MapExt(m map[string]string, keyValJon string, joinStr string, other ...string) string {
	//1 排序
	kList := make([]string, 0, len(m))
	for k := range m {
		kList = append(kList, k)
	}
	sort.Strings(kList)

	//2 按 = 号相加
	strList := make([]string, 0, len(m))
	for _, k := range kList {
		v := fmt.Sprintf("%s%s%s", k, keyValJon, m[k])
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
	r := strings.ToUpper(MD5(allStr))
	return r
}

func MD5Map(m map[string]string, joinStr string, other ...string) string {
	return MD5MapExt(m, "=", joinStr, other...)
}
