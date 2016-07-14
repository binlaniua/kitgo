package kitgo


//-------------------------------------
//
//
//
//-------------------------------------
func MapGetString(m map[string]interface{}, k string) string {
	r, ok := m[k]
	if ok {
		return r.(string)
	} else {
		return ""
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func MapGetInt64(m map[string]interface{}, k string) int64 {
	r, ok := m[k]
	if ok {
		return r.(int64)
	} else {
		return -1
	}
}
