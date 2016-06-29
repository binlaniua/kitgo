package db

import (
	"database/sql"
	"errors"
	"log"
)




//-------------------------------------
//
//
//
//-------------------------------------
func HasData(sqlStr string, args ... interface{}) bool {
	return HasDataByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func HasDataByAlias(alias string, sqlStr string, args ... interface{}) bool {
	r, err := QueryMapsByAlias(alias, sqlStr, args...)
	if err != nil {
		log.Fatal("sql 语句出错")
	}
	if len(r) > 0 {
		return true
	} else {
		return false
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMaps(sqlStr string, args ... interface{}) ([]*QueryResult, error) {
	return QueryMapsByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMapsByAlias(alias string, sqlStr string, args ... interface{}) ([]*QueryResult, error) {
	rows, err := QueryByAlias(alias, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	r := make([]map[string]*RowData, 0)
	for rows.Next() {
		rMap := mappingToMap(rows)
		r = append(r, rMap)
	}
	rr := make([]*QueryResult, 0, len(r))
	for _, v := range r {
		rr = append(rr, &QueryResult{v})
	}
	return rr, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMap(sql string, args ... interface{}) (*QueryResult, error) {
	return QueryMapByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryMapByAlias(alias string, sql string, args ... interface{}) (*QueryResult, error) {
	r, err := QueryMapsByAlias(alias, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("no data")
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