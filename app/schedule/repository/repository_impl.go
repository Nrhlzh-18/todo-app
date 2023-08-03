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

func (r *RepositoryImpl) Store(c echo.Context, db *gorm.DB, data models.MSchedule) (int64, error) {
	if err := db.Create(&data).Error; err != nil {
		return 0, err
	}

	return data.ID, nil
}

func (r *RepositoryImpl) Update(c echo.Context, db *gorm.DB, data models.MSchedule) (int64, error) {
	result := db.Model(&data).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected
	return rowsAffected, nil
}

func (r *RepositoryImpl) Delete(c echo.Context, db *gorm.DB, id string) (int64, error) {
	var schedule models.MSchedule
	if err := db.First(&schedule, id).Error; err != nil {
		return 0, err
	}

	result := db.Delete(&schedule)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected

	return rowsAffected, nil
}
