package kfile

import (
	"os"
	"github.com/binlaniua/kitgo/config"
	"io/ioutil"
	"bufio"
	"io"
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

func ReadLines(filePath string) ([]string, bool) {
	f, err := os.Open(filePath)
	if err != nil {
		config.Log(filePath, " 打开错误 =>", err)
		return nil, false
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	r := make([]string, 0)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		r = append(r, line[:len(line) - 1])
	}
	return r, true
}
