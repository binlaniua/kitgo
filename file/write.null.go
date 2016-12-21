package file

//-------------------------------------
//
//
//
//-------------------------------------
type NullWrite struct {
}

//
//
// 空写
//
//
func (n *NullWrite) Write(p []byte) (int, error) {
	return 0, nil
}
