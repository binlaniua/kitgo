package security

import (
	"crypto/des"
	"crypto/cipher"
	"bytes"
)



//-------------------------------------
//
// Des 加密
//
//-------------------------------------
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	padding := blockSize - len(origData) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	origData = append(origData, padtext...)
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	length := len(origData)
	unpadding := int(origData[length - 1])
	origData = origData[:(length - unpadding)]
	return origData, nil
}
