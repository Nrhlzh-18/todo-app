package repository

import (
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]models.MSchedule, error)
	GetById(c echo.Context, db *gorm.DB, id string) (models.MSchedule, error)
	GetByProject(c echo.Context, db *gorm.DB, projectID string) ([]models.MSchedule, error)
	GetByUser(c echo.Context, db *gorm.DB, userID string) ([]models.MSchedule, error)
	Create(c echo.Context, db *gorm.DB, data models.MTasks) (models.MSchedule, error)
	Update(c echo.Context, db *gorm.DB, data models.MSchedule) (models.MSchedule, error)
	Delete(c echo.Context, db *gorm.DB, id string) error
}
