package file

import (
	"github.com/binlaniua/kitgo"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
func RenameTo(filePath string, newName string) (string, error) {
	dir := filepath.Dir(filePath)
	newPath := dir + "/" + newName
	err := os.Rename(filePath, newPath)
	if err != nil {
		kitgo.ErrorLog.Printf("[ %s ]重命名失败 => [ %v ]", filePath, err)
		return "", err
	}
	return newPath, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func PathJoin(args ...string) string {
	s := string(os.PathSeparator)
	return strings.Join(args, s)
}

//-------------------------------------
//
//
//
//-------------------------------------
func PathToFileName(filePath string) string {
	s := strings.LastIndexAny(filePath, string(os.PathSeparator))
	if s >= 0 {
		filePath = filePath[s + 1:]
	}
	e := strings.Index(filePath, ".")
	if e >= 0 {
		filePath = filePath[:e]
	}
	return filePath
}

//-------------------------------------
//
//
//
//-------------------------------------
func ListFilePaths(dir string) ([]string, error) {
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	lf := []string{}
	for _, info := range d {
		// 完整路径
		p := dir + string(os.PathSeparator) + info.Name()

		//
		if info.IsDir() {
			// 目录
			subFileList, err := ListFilePaths(p)
			if err != nil {
				return nil, err
			} else {
				lf = append(lf, subFileList...)
			}
		} else {
			// 文件
			lf = append(lf, p)
		}
	}
	return lf, nil
}
