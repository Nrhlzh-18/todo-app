package repository

import (
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetAll(c echo.Context, db *gorm.DB) ([]models.MSchedule, error) {
	var schedule []models.MSchedule
	if err := db.Find(&schedule).Error; err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id string) (models.MSchedule, error) {
	var schedule models.MSchedule
	if err := db.First(&schedule, id).Error; err != nil {
		return schedule, err
	}

	return schedule, nil
}
