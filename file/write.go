package file

import (
	"github.com/binlaniua/kitgo"
	"os"
)

//-------------------------------------
//
//
//
//-------------------------------------
func WriteString(filePath string, src string) bool {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		kitgo.Log(filePath, "创建文件失败,", err)
		return false
	}
	defer file.Close()
	_, err = file.WriteString(src)
	if err != nil {
		kitgo.Log(filePath, "写入文件失败", err)
	}
	return true
}

//-------------------------------------
//
//
//
//-------------------------------------
func WriteBytes(filePath string, data []byte) bool {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		kitgo.Log(filePath, "创建文件失败,", err)
		return false
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		kitgo.Log(filePath, "写入文件失败", err)
	}
	return true
}
