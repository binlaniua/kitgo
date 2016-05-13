package excel

import (
	"github.com/tealeg/xlsx"
	"github.com/extrame/xls"
	"strings"
	"errors"
)

//-------------------------------------
//
//
//
//-------------------------------------
func ReadXlsxCell(filePath string, sheetIndex int, rowIndex int, cellIndex int) (string, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return "", err
	}
	if s := xlFile.Sheets[sheetIndex]; s != nil {
		if r := s.Rows[rowIndex]; r != nil {
			if c := r.Cells[cellIndex]; c != nil {
				r, e := c.String()
				return r, e
			} else {
				return "", errors.New("没有该所有的单元格")
			}
		}
	}
	return "", errors.New("没有该索引的Sheet")
}

//-------------------------------------
//
//
//
//-------------------------------------
func ReadXlsCell(filePath string, sheetIndex int, rowIndex uint16, cellIndex uint16) (string, error) {
	xlFile, err := xls.Open(filePath, "UTF-8")
	if err != nil {
		return "", err
	}
	if s := xlFile.GetSheet(sheetIndex); s != nil {
		if r := s.Rows[rowIndex]; r != nil {
			if c := r.Cols[cellIndex]; c != nil {
				return strings.Join(c.String(xlFile), ""), nil
			} else {
				return "", errors.New("没有该索引的单元格")
			}
		}
	}
	return "", errors.New("没有该索引的Sheet")
}
