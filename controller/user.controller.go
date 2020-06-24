package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/okakyo/user-api/service"
	
)
func UserRouter(e *echo.Echo){
	e.GET("/",service.FindAllUser)
	e.GET("/:userId",service.FindUserDetail)
}