package kitgo

import (
	"archive/zip"
	"io/ioutil"
	"os"
)

//-------------------------------------
//
// ZipDir zip dir files in a zip file
//
//-------------------------------------
func ZipDir(dir string, dest string) error {
	f, err := FileListPaths(dir)
	if err != nil {
		return err
	}
	return ZipFiles(f, dest)
}

//-------------------------------------
//
// ZipFiles zip provide files in a zip file
//
//-------------------------------------
func ZipFiles(fileNames []string, dest string) error {
	fzip, err := os.Create(dest)
	if err != nil {
		return err
	}
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range fileNames {
		header := &zip.FileHeader{
			Name:   file,
			Flags:  1 << 11, // 使用utf8编码
			Method: zip.Deflate,
		}
		fileZip, err := w.CreateHeader(header)
		if err != nil {
			return err
		}
		filecontent, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		_, err = fileZip.Write(filecontent)
		if err != nil {
			return err
		}
	}
	return nil
}
