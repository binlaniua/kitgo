package kitgo

//-------------------------------------
//
//
//
//-------------------------------------
func ArrayRemoveAtIndex(src []interface{}, index int) interface{} {
	if len(src) > index {
		end := index + 1
		return append(src[:index], src[end:]...)
	} else {
		return src
	}
}