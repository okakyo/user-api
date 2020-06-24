package main

import (
	"github.com/labstack/echo/v4"
	
	"github.com/okakyo/user-api/controller"
)

func main() {
  e := echo.New()


  controller.UserRouter(e)
  e.Logger.Fatal(e.Start(":1323"))
}