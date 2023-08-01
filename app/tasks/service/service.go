package service

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]tasks.TasksResponse, error)
	GetById(c echo.Context, id string) (tasks.TasksResponse, error)
	GetByProject(c echo.Context, projectID string) ([]tasks.TasksResponse, error)
	GetByUser(c echo.Context, userID string) ([]tasks.TasksResponse, error)
	Create(c echo.Context, request tasks.TaskRequest) (tasks.TasksCreateResponse, error)
	Update(c echo.Context, id string, request tasks.TaskRequest) (tasks.TasksResponse, error)
	Delete(c echo.Context, id string) error
}
