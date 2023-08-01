package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	CheckLogin(c echo.Context) error
}
