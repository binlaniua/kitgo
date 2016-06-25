package db

import (
	"database/sql"
	"strconv"
	"errors"
)

var (
	ERROR_QUERY_NO_DATA = errors.New("未查询到数据")
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

func (r *RowData) MustString() string {
	return r.ToString()
}

func (r *RowData) MustInt() int64  {
	result, _ := r.ToInt()
	return result
}

//-------------------------------------
//
//
//
//-------------------------------------
func HasData(sqlStr string, args ... interface{}) error {
	return HasDataByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func HasDataByAlias(alias string, sqlStr string, args ... interface{}) error {
	r, err := QueryMapsByAlias(alias, sqlStr, args...)
	if err != nil {
		return err
	}
	if len(r) > 0 {
		return nil
	} else {
		return ERROR_QUERY_NO_DATA
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMaps(sqlStr string, args ... interface{}) ([]map[string]*RowData, error) {
	return QueryMapsByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMapsByAlias(alias string, sqlStr string, args ... interface{}) ([]map[string]*RowData, error) {
	rows, err := QueryByAlias(alias, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	r := make([]map[string]*RowData, 0)
	for rows.Next() {
		rMap := mappingToMap(rows)
		r = append(r, rMap)
		//rows.Close()
		//log.Println(rMap)
	}
	return r, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMap(sql string, args ... interface{}) (map[string]*RowData, error) {
	return QueryMapByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMapByAlias(alias string, sql string, args ... interface{}) (map[string]*RowData, error) {
	r, err := QueryMapsByAlias(alias, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, ERROR_QUERY_NO_DATA
	}
	return r[0], nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryByAlias(alias string, sqlStr string, args ... interface{}) (*sql.Rows, error) {
	db := GetDBByAlias(alias)
	var r *sql.Rows
	var e error
	if len(args) == 0 {
		r, e = db.Query(sqlStr)
	} else {
		r, e = db.Query(sqlStr, args...)
	}
	return r, e
}

func Query(sqlStr string, args ... interface{}) (*sql.Rows, error) {
	return QueryByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func mappingToMap(row *sql.Rows) map[string]*RowData {
	columns, _ := row.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	row.Scan(scanArgs...)
	r := make(map[string]*RowData)
	for i, col := range values {
		buff := make([]byte, 0, len(col))
		buff = append(buff, col...)
		r[columns[i]] = &RowData{buff}
	}
	return r
}