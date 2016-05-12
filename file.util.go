package kitgo

import (
	"os"
	"path/filepath"
)



//-------------------------------------
//
//
//
//-------------------------------------
func FileRenameTo(filePath string, newName string) (string, bool) {
	dir := filepath.Dir(filePath)
	newPath := dir + "/" + newName
	err := os.Rename(filePath, newPath)
	if err != nil {
		Log(filePath, " 重命名失败 => ", err)
		return "", false
	}
	return newPath, true
}
