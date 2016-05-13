package file

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
func ReadBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadString(filePath string) (string, error) {
	data, err := ReadBytes(filePath)
	if err != nil {
		return string(data), err
	} else {
		return "", nil
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
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
	return r, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	} else {
		return true
	}
}

//-------------------------------------
//
//  读取Json文件
//
//-------------------------------------
func LoadJsonFile(filePath string, obj interface{}) error {
	data, err := ReadBytes(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

//-------------------------------------
//
//  读取配置文件
//
//-------------------------------------
func LoadProperties(filePath string) bool {
	return false
}
