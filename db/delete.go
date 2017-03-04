package db

//-------------------------------------
//
//
//
//-------------------------------------
func Delete(sql string, args ...interface{}) (int64, error) {
	return DeleteByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func DeleteByAlias(alias string, sql string, args ...interface{}) (int64, error) {
	r, err := DMLByAlias(alias, sql, args...)
	if err != nil {
		errorLogger.Printf("删除[ %s ][ %v ] => [ %v ]", sql, args, err)
		return 0, err
	} else {
		debugLogger.Printf("删除[ %s ] => [ %v ]", sql, r)
		id, _ := r.LastInsertId()
		return id, nil
	}
}
