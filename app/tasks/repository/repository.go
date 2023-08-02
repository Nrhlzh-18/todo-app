package repository

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]tasks.TasksResponse, error)
	GetById(c echo.Context, db *gorm.DB, id string) (tasks.TasksResponse, error)
	Create(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error)
	Update(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error)
	Delete(c echo.Context, db *gorm.DB, id string) error
}
