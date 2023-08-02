package routes

import (
	tasks "github.com/Nrhlzh-18/todo-app/app/tasks/handler"
	user "github.com/Nrhlzh-18/todo-app/app/user/handler"
	"github.com/Nrhlzh-18/todo-app/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoutes(e *echo.Echo, db *gorm.DB) {
	// Routes
	g := e.Group("/api/v1")

	tasks.NewHandler(db).Route(g.Group("/tasks"))
	user.NewHandler(db).Route(g.Group("/user"))

	e.HTTPErrorHandler = exception.NewErrorHandler
}
