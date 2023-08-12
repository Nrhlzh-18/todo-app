package routes

import (
	auth "github.com/Nrhlzh-18/todo-app/app/auth/handler"
	schedule "github.com/Nrhlzh-18/todo-app/app/schedule/handler"
	tasks "github.com/Nrhlzh-18/todo-app/app/tasks/handler"
	user "github.com/Nrhlzh-18/todo-app/app/user/handler"

	"github.com/Nrhlzh-18/todo-app/exception"
	authentication "github.com/Nrhlzh-18/todo-app/helpers/is_authenticated"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoutes(e *echo.Echo, db *gorm.DB) {
	// Routes
	g := e.Group("/api/v1")
	a := e.Group("/api/v1")

	a.Use(authentication.IsAuthenticated)

	tasks.NewHandler(db).Route(a.Group("/tasks"))
	user.NewHandler(db).Route(a.Group("/user"))
	schedule.NewHandler(db).Route(a.Group("/schedule"))
	auth.NewHandler(db).Route(g.Group(""))

	e.HTTPErrorHandler = exception.NewErrorHandler
}
