package db

//-------------------------------------
//
//
//
//-------------------------------------
func Update(sql string, args ...interface{}) (int64, error) {
	return UpdateByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func UpdateByAlias(alias string, sql string, args ...interface{}) (int64, error) {
	r, err := DMLByAlias(alias, sql, args...)
	if err != nil {
		errorLogger.Printf("更新[ %s ][ %v ] => [ %v ]", sql, args, err)
		return 0, err
	} else {
		debugLogger.Printf("更新[ %s ] => [ %v ]", sql, r)
		id, _ := r.LastInsertId()
		return id, nil
	}
}
