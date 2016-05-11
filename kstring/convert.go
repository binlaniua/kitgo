package kstring

import (
	"strconv"
	"github.com/binlaniua/kitgo/kconfig"
)

//-------------------------------------
//
//
//
//-------------------------------------
func FromUnicode(src string) (string, bool) {
	str, err := strconv.Unquote(`"` + src + `"`)
	if err != nil {
		kconfig.Log(src, " 转换到中文失败 => ", err)
	}
	return str, err != nil
}
