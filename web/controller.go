package web

import "github.com/labstack/echo"

//-------------------------------------
//
//
//
//-------------------------------------
type Controller interface {

	//
	//
	// 初始化
	//
	//
	Init(echo *echo.Echo)
}
