package kitgo

import (
	"os"
	"io/ioutil"
	"bufio"
	"io"
	"encoding/json"
)

//-------------------------------------
//
//
//
//-------------------------------------
func FileReadBytes(filePath string) ([]byte, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		Log(filePath, "文件不存在, 无法读取")
		return nil, false
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		Log(filePath, "读取文件出错, ", err)
		return nil, false
	}
	return data, true
}

//-------------------------------------
//
//
//
//-------------------------------------
func FileReadString(filePath string) (string, bool) {
	data, ok := FileReadBytes(filePath)
	if ok {
		return string(data), true
	} else {
		return "", false
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func FileReadLines(filePath string) ([]string, bool) {
	f, err := os.Open(filePath)
	if err != nil {
		Log(filePath, " 打开错误 =>", err)
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

//-------------------------------------
//
//  读取Json文件
//
//-------------------------------------
func FileLoadJsonFile(filePath string, obj interface{}) bool {
	data, ok := FileReadBytes(filePath)
	if !ok {
		return false
	}
	err := json.Unmarshal(data, obj)
	if err != nil {
		Log(filePath, "加载失败 => ", err)
		return false
	}
	return true
}

//-------------------------------------
//
//  读取配置文件
//
//-------------------------------------
func FileLoadProperties(filePath string) bool {
	return false
}
