package excel

import (
	"reflect"
	"github.com/tealeg/xlsx"
	"os"
)

//-------------------------------------
//
//  dataList 里面的字段必须有 excel tag
//
//-------------------------------------
func WriteXlsx(filePath string, dataList interface{}) error {
	values := reflect.ValueOf(dataList)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("0")
	headRow := sheet.AddRow()
	if err != nil {
		return err
	}
	for i := 0; i < values.Len(); i++ {
		dataRow := sheet.AddRow()
		item := values.Index(i)
		if i == 0 {
			writeDataToRow(headRow, item, true)
		}
		writeDataToRow(dataRow, item, false)
	}
	xlsFile.Write(file)
	return nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func writeDataToRow(row *xlsx.Row, data reflect.Value, isHead bool) {
	dataItemType := data.Type()
	for j := 0; j < dataItemType.NumField(); j++ {
		field := dataItemType.Field(j)
		if columnName := field.Tag.Get("excel"); columnName != "" {
			if isHead {
				headCell := row.AddCell()
				headCell.SetString(columnName)
			} else {
				dataStr := data.Field(j).String()
				dataCell := row.AddCell()
				dataCell.SetString(dataStr)
			}
		}
	}
}