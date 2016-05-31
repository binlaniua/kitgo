package zip

import (
	"io/ioutil"
	"archive/zip"
	"os"
)

//-------------------------------------
//
//
//
//-------------------------------------
func Dir(dir string, dist string) error {
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	fzip, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer fzip.Close()
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		header := &zip.FileHeader{
			Name:   file.Name(),
			Flags:  1 << 11, // 使用utf8编码
			Method: zip.Deflate,
		}
		fileZip, err := w.CreateHeader(header)
		if err != nil {
			return err
		}
		filecontent, err := ioutil.ReadFile(dir + file.Name())
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
