package repository

import (
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]models.MSchedule, error)
	GetById(c echo.Context, db *gorm.DB, id string) (models.MSchedule, error)
	GetByDate(c echo.Context, db *gorm.DB) ([]models.MSchedule, error)
	Store(c echo.Context, db *gorm.DB, data *models.MSchedule) error
	Update(c echo.Context, db *gorm.DB, data *models.MSchedule) (int64, error)
	Delete(c echo.Context, db *gorm.DB, id string) (int64, error)
}
