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
		return 0, err
	} else {
		id, _ := r.LastInsertId()
		return id, nil
	}
}
