package repository

import (
	taskstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(c echo.Context, db *gorm.DB, projectID string) (taskstag.TasksTagProject, error)
	GetByProject(c echo.Context, db *gorm.DB, projectID string) ([]taskstag.TasksTagProject, error)
	GetByTeam(c echo.Context, db *gorm.DB, teamID string) ([]taskstag.TasksTagTeam, error)
	GetByUser(c echo.Context, db *gorm.DB, UserID string) ([]taskstag.TasksTagUser, error)
	Store(c echo.Context, db *gorm.DB, data models.TasksTag) (int64, error)
}

