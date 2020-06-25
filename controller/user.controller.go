package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/okakyo/user-api/service"
	
)
func UserRouter(e *echo.Echo){
	g:=e.Group("/user")
	g.GET("/",service.FindAllUser)
	g.GET("/:userId",service.FindUserDetail)
}