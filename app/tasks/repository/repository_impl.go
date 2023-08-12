package repository

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetAll(c echo.Context, db *gorm.DB) ([]tasks.TasksResponse, error) {
	var taskList []tasks.TasksResponse
	if err := db.Find(&taskList).Error; err != nil {
		return taskList, err
	}

	return taskList, nil
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id string) (tasks.TasksResponse, error) {
	var taskData tasks.TasksResponse
	if err := db.Where("id = ?", id).First(&taskData).Error; err != nil {
		return taskData, err
	}

	return taskData, nil
}

func (r *RepositoryImpl) ChangePriority(c echo.Context, db *gorm.DB, priority, id string) error {
	var existingData models.MTasks
	if err := db.Model(&existingData).Where("id = ?", id).Update("priority", priority).Error; err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) ChangeStatus(c echo.Context, db *gorm.DB, status, id string) error {
	var existingData models.MTasks
	if err := db.Model(&existingData).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) Create(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error) {
	if err := db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *RepositoryImpl) Update(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error) {
	var existingData models.MTasks

	if err := db.First(&existingData, data.ID).Error; err != nil {
		return data, err
	}

	if err := db.Model(&existingData).Updates(data).Error; err != nil {
		return data, err
	}

	return existingData, nil
}

func (r *RepositoryImpl) Delete(c echo.Context, db *gorm.DB, id string) error {
	var existingData models.MTasks

	if err := db.First(&existingData, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&existingData).Error; err != nil {
		return err
	}

	return nil
}
