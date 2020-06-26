package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/okakyo/user-api/handler/controller"
)

func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())


  controller.UserRouter(e)
  e.Logger.Fatal(e.Start(":1323"))
}