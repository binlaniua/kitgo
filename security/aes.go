package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//-------------------------------------
//
// 加密
//
//-------------------------------------
func AESEncrypt(key []byte, src string) ([]byte, error) {
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(src))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(src))
	return encrypted, nil
}

//-------------------------------------
//
// 加密
//
//-------------------------------------
func AESDecrypt(key []byte, src []byte) (strDesc string, err error) {
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
