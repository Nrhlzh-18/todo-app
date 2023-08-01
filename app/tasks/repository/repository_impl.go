package repository

import (
	"fmt"

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

	fmt.Println(taskList)

	return taskList, nil
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id string) (tasks.TasksResponse, error) {
	var taskData tasks.TasksResponse
	if err := db.Where("id = ?", id).First(&taskData).Error; err != nil {
		return taskData, err
	}

	fmt.Println(taskData)

	return taskData, nil
}

func (r *RepositoryImpl) GetByProject(c echo.Context, db *gorm.DB, projectID string) ([]tasks.TasksResponse, error) {
	var taskList []tasks.TasksResponse
	if err := db.Where("project_id = ?", projectID).Find(&taskList).Error; err != nil {
		return taskList, err
	}

	fmt.Println(taskList)

	return taskList, nil
}

func (r *RepositoryImpl) GetByUser(c echo.Context, db *gorm.DB, userID string) ([]tasks.TasksResponse, error) {
	var taskList []tasks.TasksResponse
	if err := db.Where("user_id = ?", userID).Find(&taskList).Error; err != nil {
		return taskList, err
	}

	fmt.Println(taskList)

	return taskList, nil
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
