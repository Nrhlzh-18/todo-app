package controller

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	GetByProject(c echo.Context) error
	GetByTeam(c echo.Context) error
	GetByUser(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	UpdateStatus(c echo.Context) error
	UpdatePriority(c echo.Context) error
	Delete(c echo.Context) error
	CreateTags(c echo.Context) error
}
