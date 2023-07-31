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

func (r *RepositoryImpl) GetAll(c echo.Context, db *gorm.DB) ([]models.MTasks, error) {
	var list_tasks []models.MTasks
	if err := db.Order("created_at desc").Find(&list_tasks).Error; err != nil {
		return list_tasks, err
	}

	return list_tasks, nil
}
