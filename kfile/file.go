package kfile

import (
	"os"
	"path/filepath"
	"os/exec"
	"github.com/binlaniua/kitgo/kconfig"
)

//-------------------------------------
//
//
//
//-------------------------------------
func RuntimePath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

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
		kconfig.Log(filePath, " 重命名失败 => ", err)
		return "", false
	}
	return newPath, true
}
