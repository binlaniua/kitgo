package kitgo

import (
	"os"
)

//-------------------------------------
//
//
//
//-------------------------------------
func FileWriteString(filePath string, src string) bool {
	return FileWriteBytes(filePath, []byte(src));
}

//-------------------------------------
//
//
//
//-------------------------------------
func FileWriteBytes(filePath string, data []byte) bool {
	file, err := os.Open(filePath)
	if err != nil {
		//kconfig.Log(filePath, "打开文件失败,", err)
		file, err = os.Create(filePath)
		if err != nil {
			Log(filePath, "创建文件失败,", err)
			return false
		}
	}
	defer file.Close()
	file.Write(data)
	return true
}
