package kitgo

import (
	"strconv"
	"github.com/axgle/mahonia"
	"math"
	"fmt"
	"bytes"
	"encoding/binary"
)

var (
	decodeGBToUTF = mahonia.NewDecoder("GB18030")
)

//-------------------------------------
//
//
//
//-------------------------------------
func ConvertGBToUTF(src string) string {
	return decodeGBToUTF.ConvertString(src)
}

//-------------------------------------
//
//
//
//-------------------------------------
func ConvertFromUnicode(src string) (string, bool) {
	str, err := strconv.Unquote(`"` + src + `"`)
	if err != nil {
		Log(src, " 转换到中文失败 => ", err)
		return "", false
	}
	return str, true
}

//-------------------------------------
//
// 
//
//-------------------------------------
func ConvertToInt(src string) (int, bool) {
	if src == "" {
		return 0, false
	} else {
		r, err := strconv.Atoi(src)
		if err != nil {
			return 0, false
		} else {
			return r, true
		}
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func ConvertMustInt(src string) int {
	r, _ := ConvertToInt(src)
	return r
}


//-------------------------------------
//
//
//
//-------------------------------------
func ConvertToFloat(src string) (float64, error) {
	r, err := strconv.ParseFloat(src, 64)
	return r, err
}

//-------------------------------------
//
//
//
//-------------------------------------
func ConvertMustFloat(src string) float64 {
	r, _ := ConvertToFloat(src)
	return r
}

//-------------------------------------
//
//
//
//-------------------------------------
func ToFixed(src float64, n int) float64 {
	o := math.Pow(10, float64(n))
	r := float64(int(src * o)) / o
	return r
}

//-------------------------------------
//
//  int 转 二进制
//
//-------------------------------------
func IntToBinary(n int) string {
	return fmt.Printf("%08b", n)
}

//-------------------------------------
//
// 二进制转 int
//
//-------------------------------------
func BinaryToInt(src string) int {
	r, _ := strconv.ParseInt(src, 2, 32)
	return r
}

//-------------------------------------
//
// 高位
//
//-------------------------------------
func IntToBytes(n int) []byte {
	m := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, m)
	return bytesBuffer.Bytes()
}

//-------------------------------------
//
//  高位
//
//-------------------------------------
func BytesToInt(src []byte) int {
	bytesBuffer := bytes.NewBuffer(src)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}