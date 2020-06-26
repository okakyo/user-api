package service

import (
	"net/http"
	"github.com/labstack/echo/v4"
)


/*
- User 情報についての制御についての関数（CRUD）

*/


func FindAllUser(c echo.Context) error {
		return c.String(http.StatusOK,"Hello World")
	}


func FindUserDetail(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Goog Luck!"+userId)
	}


func createUser(c echo.Context) error {
		return c.String(http.StatusOK,"Good Morning!")
	}


func updateUser(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Welcome "+userId)
	}

func deleteUser(c echo.Context) error {
		userId:= c.Param("userId")
		return c.String(http.StatusOK,"Goog bye! "+userId)
	}

