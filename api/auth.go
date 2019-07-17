package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (db *MongoDB) LogIn(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) Register(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) LogOut(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}
