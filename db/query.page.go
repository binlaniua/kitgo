package db

//
//
//
//
//
type Pageable struct {
	Index  int
	Size   int

	Result interface{}
	Total  int64
}
