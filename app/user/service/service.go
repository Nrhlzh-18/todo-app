package service

import (
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]user.UserResponse, error)
	GetById(c echo.Context, id string) (user.UserResponse, error)
	Create(c echo.Context, request user.UserRequest) (user.UserCreateResponse, error)
	Update(c echo.Context, id string, request user.UserRequest) (user.UserResponse, error)
	Delete(c echo.Context, id string) error
}
