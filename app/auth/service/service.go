package service

import (
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/labstack/echo/v4"
)

type Service interface {
	CheckLogin(c echo.Context, user user.UserLogin) (string, error)
}
