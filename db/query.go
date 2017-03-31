package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	dbFieldMap = map[reflect.Type]map[string]reflect.StructField{}
)

//-------------------------------------
//
//
//
//-------------------------------------
func HasData(sqlStr string, args ...interface{}) bool {
	return HasDataByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func HasDataByAlias(alias string, sqlStr string, args ...interface{}) bool {
	r, err := QueryMapsByAlias(alias, sqlStr, args...)
	if err != nil {
		errorLogger.Printf("[ %s ] 语句出错 => [ %v ]", sqlStr, err)
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
func QueryMaps(sqlStr string, args ...interface{}) ([]*QueryResult, error) {
	return QueryMapsByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

func QueryMapsByAlias(alias string, sqlStr string, args ...interface{}) ([]*QueryResult, error) {
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
func QueryMap(sql string, args ...interface{}) (*QueryResult, error) {
	return QueryMapByAlias(DEFAULT_DB_NAME, sql, args...)
}

func QueryMapByAlias(alias string, sql string, args ...interface{}) (*QueryResult, error) {
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
func QueryByAlias(alias string, sqlStr string, args ...interface{}) (*sql.Rows, error) {
	db := GetDBByAlias(alias)
	var r *sql.Rows
	var e error
	if len(args) == 0 || args == nil {
		r, e = db.Query(sqlStr)
	} else {
		r, e = db.Query(sqlStr, args...)
	}
	if e != nil {
		errorLogger.Printf("查询[ %s ][ %v ] => [ %v ]", sqlStr, args, e)
		return nil, e
	} else {
		debugLogger.Printf("查询[ %s ] => ok", sqlStr)
	}
	return r, e
}

func Query(sqlStr string, args ...interface{}) (*sql.Rows, error) {
	return QueryByAlias(DEFAULT_DB_NAME, sqlStr, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryObjectByAlias(alias string, sqlStr string, obj interface{}, args ...interface{}) error {
	db := GetDBByAlias(alias)
	var r *sql.Rows
	var e error
	if len(args) == 0 || args == nil {
		r, e = db.Query(sqlStr)
	} else {
		r, e = db.Query(sqlStr, args...)
	}
	if e != nil {
		errorLogger.Printf("查询[ %s ][ %v ] => [ %v ]", sqlStr, args, e)
		return e
	} else {
		debugLogger.Printf("查询[ %s ] => ok", sqlStr)
	}
	defer r.Close()
	r.Next()
	ov := reflect.ValueOf(obj)
	mappingToObject(r, ov)
	return nil
}

func QueryObject(sqlStr string, obj interface{}, args ...interface{}) error {
	return QueryObjectByAlias(DEFAULT_DB_NAME, sqlStr, obj, args...)
}

//
//
// 分页查询
//
//
func QueryPage(sqlStr string, page *Pageable, args ...interface{}) error {
	return QueryPageByAlias(DEFAULT_DB_NAME, sqlStr, page, args...)
}

func QueryPageByAlias(alias string, sqlStr string, pageable *Pageable, args ...interface{}) error {
	//1 先查询数量
	countSqlStr := fmt.Sprintf("SELECT count(1) as count from (%s) t", sqlStr)
	m, err := QueryMapByAlias(alias, countSqlStr, args...)
	if err != nil {
		return err
	}
	total := m.GetInt64("count")
	pageable.Total = total

	//2. 判断够不够
	if pageable.Index < 0 {
		pageable.Index = 0
	}
	start := int64(pageable.Index * pageable.Size)
	if total < start {
		return nil
	}
	limitSqlStr := fmt.Sprintf("select __t.* from (%s) __t limit %d, %d", sqlStr, start, pageable.Size)
	err = QueryListByAlias(alias, limitSqlStr, pageable.Result, args...)
	return err
}

//-------------------------------------
//
//
//
//-------------------------------------
func QueryList(sqlStr string, result interface{}, args ...interface{}) error {
	return QueryListByAlias(DEFAULT_DB_NAME, sqlStr, result, args...)
}

func QueryListByAlias(alias string, sqlStr string, result interface{}, args ...interface{}) error {
	db := GetDBByAlias(alias)
	var r *sql.Rows
	var e error

	//
	resultList := reflect.Indirect(reflect.ValueOf(result))
	resultElementType := resultList.Type().Elem().Elem()
	if len(args) == 0 || args == nil {
		r, e = db.Query(sqlStr)
	} else {
		r, e = db.Query(sqlStr, args...)
	}

	//
	if e != nil {
		errorLogger.Printf("查询[ %s ][ %v ] => [ %v ]", sqlStr, args, e)
		return e
	} else {
		debugLogger.Printf("查询[ %s ] => ok", sqlStr)
	}
	defer r.Close()

	//
	for r.Next() {
		newValue := reflect.New(resultElementType)
		mappingToObject(r, newValue)
		resultList.Set(reflect.Append(resultList, newValue))
	}
	return nil
}

//-------------------------------------
//
//
//
//-------------------------------------
var (
	timeType = reflect.TypeOf(time.Time{})
)

func mappingToObject(row *sql.Rows, newValue reflect.Value) {
	fieldMap := mappingFieldMap(newValue.Elem().Type())
	columns, _ := row.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	row.Scan(scanArgs...)
	for i, col := range values {
		field, ok := fieldMap[strings.ToLower(columns[i])]
		if !ok {
			continue
		} else {
			rowData := &RowData{col}
			valueField := newValue.Elem().FieldByName(field.Name)
			switch field.Type.Kind() {
			case reflect.Int:
				r, _ := rowData.ToInt32()
				valueField.Set(reflect.ValueOf(r))
			case reflect.Int64:
				r, _ := rowData.ToInt()
				valueField.Set(reflect.ValueOf(r))
			case reflect.String:
				valueField.Set(reflect.ValueOf(rowData.ToString()))
			default:
				if field.Type.ConvertibleTo(timeType) {
					ts := rowData.ToString()
					t, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
					valueField.Set(reflect.ValueOf(t))
				} else {
					errorLogger.Printf("映射[ %v ]没有匹配的 => [ %v ]", field.Type.Kind(), field.Type)
				}
			}
		}
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
var lock = sync.RWMutex{}

func mappingFieldMap(class reflect.Type) map[string]reflect.StructField {
	m, ok := dbFieldMap[class]
	if ok {
		return m
	} else {
		lock.Lock()
		defer lock.Unlock()
		m := map[string]reflect.StructField{}
		numCount := class.NumField()
		for i := 0; i < numCount; i++ {
			field := class.Field(i)
			dbName := strings.ToLower(field.Tag.Get("db"))
			m[dbName] = field
		}
		dbFieldMap[class] = m
		return m
	}
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
