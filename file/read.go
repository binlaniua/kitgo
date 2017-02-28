package file

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
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
		return "", err
	} else {
		return string(data), err
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadLines(filePath string) ([]string, error) {
	r := []string{}
	e := ReadLinesExt(filePath, func(line string, row int64) bool {
		r = append(r, line)
		return true
	})
	return r, e
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadLinesExt(filePath string, fn func(string, int64) bool) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	row := 0
	for {
		row++
		line, err := rd.ReadString('\n')
		if err != nil && io.EOF == err {
			return nil
		} else if err != nil {
			return err
		}
		line = line[:len(line)-1]
		r := fn(line, int64(row))
		if !r {
			break
		}
	}
	return nil
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
//
//
//-------------------------------------
func LoadJsonFileToMap(filePath string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	e := LoadJsonFile(filePath, &m)
	return m, e
}
