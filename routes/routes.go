package routes

import (
	"github.com/Nrhlzh-18/todo-app/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Routes struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewRoutes(e *echo.Echo, db *gorm.DB) *Routes {
	return &Routes{
		Echo: e,
		DB:   db,
	}
}

func (r *Routes) Register() {
	// Routes
	_ = r.Echo.Group("/api/v1")


	r.Echo.HTTPErrorHandler = exception.NewErrorHandler
}
