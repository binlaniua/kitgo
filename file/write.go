package file

import (
	"os"
	"github.com/binlaniua/kitgo/config"
)

//-------------------------------------
//
//
//
//-------------------------------------
func WriteString(filePath string, src string) bool {
	return WriteBytes(filePath, []byte(src));
}

//-------------------------------------
//
//
//
//-------------------------------------
func WriteBytes(filePath string, data []byte) bool {
	file, err := os.Open(filePath)
	if err != nil {
		//kconfig.Log(filePath, "打开文件失败,", err)
		file, err = os.Create(filePath)
		if err != nil {
			config.Log(filePath, "创建文件失败,", err)
			return false
		}
	}
	defer file.Close()
	file.Write(data)
	return true
}
