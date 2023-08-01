package repository

import (
	"fmt"

	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetAll(c echo.Context, db *gorm.DB) ([]user.UserResponse, error) {
	var taskList []user.UserResponse
	if err := db.Find(&taskList).Error; err != nil {
		return taskList, err
	}

	fmt.Println(taskList)

	return taskList, nil
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id int64) (user.UserResponse, error) {
	var user user.UserResponse
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *RepositoryImpl) GetByUsernamePass(c echo.Context, db *gorm.DB, user *user.UserLogin) (models.MUser, error) {
	var users models.MUser
	if err := db.Where("username = ?", user.Username).First(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *RepositoryImpl) Create(c echo.Context, db *gorm.DB, data models.MUser) (models.MUser, error) {
	if err := db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *RepositoryImpl) Update(c echo.Context, db *gorm.DB, data models.MUser) (models.MUser, error) {
	var existingData models.MUser

	if err := db.First(&existingData, data.ID).Error; err != nil {
		return data, err
	}

	if err := db.Model(&existingData).Updates(data).Error; err != nil {
		return data, err
	}

	return existingData, nil
}

func (r *RepositoryImpl) Delete(c echo.Context, db *gorm.DB, id int64) error {

	var data models.MUser
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return err
	}

	result := db.Delete(&data, "id=?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
