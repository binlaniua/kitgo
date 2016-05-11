package kexcel

import (
	"github.com/tealeg/xlsx"
	"github.com/extrame/xls"
	"github.com/binlaniua/kitgo/kconfig"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
func ReadXlsxCell(filePath string, sheetIndex int, rowIndex int, cellIndex int) (string, bool) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		kconfig.Log(filePath, " 读取excel报错 => ", err)
		return "", false
	}
	if s := xlFile.Sheets[sheetIndex]; s != nil {
		if r := s.Rows[rowIndex]; r != nil {
			if c := r.Cells[cellIndex]; c != nil {
				r, e := c.String()
				return r, e == nil
			} else {
				return "", false
			}
		}
	}
	return "", false
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadXlsCell(filePath string, sheetIndex int, rowIndex uint16, cellIndex uint16) (string, bool) {
	xlFile, err := xls.Open(filePath, "UTF-8")
	if err != nil {
		kconfig.Log(filePath, " 读取xls出错 => ", err)
		return "", false
	}
	if s := xlFile.GetSheet(sheetIndex); s != nil {
		if r := s.Rows[rowIndex]; r != nil {
			if c := r.Cols[cellIndex]; c != nil {
				return strings.Join(c.String(xlFile), ""), true
			} else {
				return "", false
			}
		}
	}
	return "", false
}
