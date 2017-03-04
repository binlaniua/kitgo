package random

import (
	"math/rand"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
func do(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			// random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

//-------------------------------------
//
//
//
//-------------------------------------
func RandomInt(max int) int {
	return rand.Intn(max)
}

//-------------------------------------
//
//
//
//-------------------------------------
func RandomNum(size int) string {
	return do(size, 0)
}

//-------------------------------------
//
//
//
//-------------------------------------
func RandomLowString(size int) string {
	return do(size, 1)
}

//-------------------------------------
//
//
//
//-------------------------------------
func RandomUpperString(size int) string {
	return do(size, 2)
}

//-------------------------------------
//
//
//
//-------------------------------------
func RandomString(size int) string {
	return do(size, 3)
}
