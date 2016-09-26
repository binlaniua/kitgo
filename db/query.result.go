package db

import (
	"database/sql"
	"strconv"
)


//-------------------------------------
//
//
//
//-------------------------------------
type RowData struct {
	data sql.RawBytes
}

func (r *RowData) ToString() string {
	return string(r.data)
}

func (r *RowData) ToInt() (int64, error) {
	s := r.ToString()
	re, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return re, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
type QueryResult struct {
	dataMap map[string]*RowData
}

func (r *QueryResult) ConvertTo(target interface{}) bool {
	return true
}

func (r *QueryResult) GetString(key string) string {
	rr, ok := r.dataMap[key]
	if ok {
		return rr.ToString()
	} else {
		return ""
	}
}

func (r *QueryResult) GetInt64(key string) int64 {
	rr, ok := r.dataMap[key]
	if ok {
		result, err := rr.ToInt()
		if err != nil {
			return -1
		} else {
			return result
		}
	} else {
		return -1
	}
}

func (r *QueryResult) GetInt(key string) int {
	return int(r.GetInt64(key))
}

func (r *QueryResult) GetDataMap() map[string]*RowData {
	return r.dataMap
}