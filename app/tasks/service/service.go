package service

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]tasks.TasksResponse, error)
	// Create(c echo.Context, request tasks.TaskRequest) (tasks.TasksResponse, error)
	// Update(c echo.Context, request tasks.TaskRequest) (tasks.TasksResponse, error)
	// Delete(c echo.Context, refId uint64) error
}
