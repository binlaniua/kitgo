package file

import (
	"os"
	"path/filepath"
	"github.com/binlaniua/kitgo"
	"io/ioutil"
	"strings"
)



//-------------------------------------
//
//
//
//-------------------------------------
func RenameTo(filePath string, newName string) (string, bool) {
	dir := filepath.Dir(filePath)
	newPath := dir + "/" + newName
	err := os.Rename(filePath, newPath)
	if err != nil {
		kitgo.Log(filePath, " 重命名失败 => ", err)
		return "", false
	}
	return newPath, true
}

//-------------------------------------
//
//
//
//-------------------------------------
func PathJoin(args ... string) string {
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
		if info.IsDir() {
			continue
		}
		p := dir + string(os.PathSeparator) + info.Name()
		lf = append(lf, p)
	}
	return lf, nil
}