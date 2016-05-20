package security

import (
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
	"crypto/sha1"
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
func SHA1(s string) string {
	t := sha1.New();
	io.WriteString(t, s);
	return hex.EncodeToString(t.Sum(nil))
}
