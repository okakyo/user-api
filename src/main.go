package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"okakyo/user-api/src/controller"
)

func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  controller.userRouter(e)
  e.Logger.Fatal(e.Start(":1323"))
}