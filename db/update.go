package db



//-------------------------------------
//
//
//
//-------------------------------------
func Update(sql string, args ... interface{}) (int64, error) {
	return UpdateByAlias(DEFAULT_DB_NAME, sql, args...)
}

//-------------------------------------
//
//
//
//-------------------------------------
func UpdateByAlias(alias string, sql string, args ... interface{}) (int64, error) {
	r, err := DMLByAlias(alias, sql, args...)
	if err != nil {
		return 0, err
	} else {
		id, _ := r.LastInsertId()
		return id, nil
	}
}