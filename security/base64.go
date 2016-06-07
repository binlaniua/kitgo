package security

import "encoding/base64"

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var (
	coder = base64.NewEncoding(base64Table)
)

//-------------------------------------
//
//
//
//-------------------------------------
func ToBase64(data []byte) []byte {
	return []byte(coder.EncodeToString(data))
}

//-------------------------------------
//
//
//
//-------------------------------------
func FromBase64(data []byte) ([]byte, error) {
	return coder.DecodeString(string(data))
}
