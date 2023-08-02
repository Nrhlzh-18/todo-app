package service

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	taskstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]tasks.TasksResponse, error)
	GetById(c echo.Context, id string) (tasks.TasksResponse, error)
	GetByProject(c echo.Context, projectID string) ([]taskstag.TasksTagProject, error)
	GetByTeam(c echo.Context, teamID string) ([]taskstag.TasksTagTeam, error)
	GetByUser(c echo.Context, userID string) ([]taskstag.TasksTagUser, error)
	Create(c echo.Context, request tasks.TaskRequest) (tasks.TasksCreateResponse, error)
	Update(c echo.Context, id string, request tasks.TaskRequest) (tasks.TasksResponse, error)
	Delete(c echo.Context, id string) error
}
