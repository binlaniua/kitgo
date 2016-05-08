package kfile

import (
	"os"
	"github.com/binlaniua/kitgo/config"
	"io/ioutil"
)

func ReadBytes(filePath string) ([]byte, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		config.Log(filePath, "文件不存在, 无法读取")
		return nil, false
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		config.Log(filePath, "读取文件出错, ", err)
		return nil, false
	}
	return data, true
}

func ReadString(filePath string) (string, bool) {
	data, ok := ReadBytes(filePath)
	if ok {
		return string(data), true
	} else {
		return "", false
	}
}

func ReadStringIE(filePath string) string  {
	r, _ := ReadString(filePath)
	return r
}
