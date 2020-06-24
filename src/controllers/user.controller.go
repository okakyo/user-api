package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/okakyo/user-api/src/service"
	
)
func userRouter(e *echo.Echo){
	e.GET("/",service.findAllUser)
	e.GET("/:id",service.findUserDetail)
}