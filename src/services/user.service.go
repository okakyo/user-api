package service

import (
	"net/http"
	"github.com/labstack/echo"
)


/*
- User 情報についての制御についての関数（CRUD）

*/


func findAllUser() echo.HandlerFunc{
	return func (c echo.Context) error {
		return c.String(http.StatusOK,"Hello World")
	}
}

func findUserDetail() echo.HandlerFunc{
	return func(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Goog Luck! "+userId)
	}
}

func createUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK,"Good Morning!")
	}
}

func updateUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Welcome "+userId)
	}
}
func deleteUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Goog bye! "+userId)
	}
}
