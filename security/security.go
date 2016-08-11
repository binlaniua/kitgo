package security

import (
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




