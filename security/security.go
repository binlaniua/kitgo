package security

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
)

//-------------------------------------
//
// 
//
//-------------------------------------
func Guid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return MD5(base64.URLEncoding.EncodeToString(b))
}

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
