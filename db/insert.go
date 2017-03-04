package db

//-------------------------------------
//
//
//
//-------------------------------------
func Insert(sql string, args ...interface{}) (int64, error) {
	return InsertByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func InsertByAlias(alias string, sql string, args ...interface{}) (int64, error) {
	r, err := DMLByAlias(alias, sql, args...)
	if err != nil {
		errorLogger.Printf("新增[ %s ][ %v ] => [ %v ]", sql, args, err)
		return 0, err
	} else {
		id, _ := r.LastInsertId()
		debugLogger.Printf("删除[ %s ] => [ %v ]", sql, id)
		return id, nil
	}
}
