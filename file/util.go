package file

import (
	"os"
	"path/filepath"
	"github.com/binlaniua/kitgo"
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