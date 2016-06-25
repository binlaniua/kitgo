package security

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
	"fmt"
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
	return MD5(s + salt);
}

//-------------------------------------
//
// 有序MD5加密
//
//-------------------------------------
func MD5Map(m map[string]string, other ... string) string {
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
	allStr := strings.Join(strList, "&")
	if other != nil {
		for _, otherItem := range other  {
			allStr = allStr + otherItem
		}
	}

	//4. 大写加密
	r := strings.ToUpper(MD5(allStr))
	return r
}

