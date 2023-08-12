package repository

import (
	"fmt"

	"github.com/Nrhlzh-18/todo-app/app/role"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetAll(c echo.Context, db *gorm.DB) ([]role.RoleResponse, error) {
	var role []role.RoleResponse
	if err := db.Find(&role).Error; err != nil {
		return role, err
	}

	fmt.Println(role)

	return role, nil
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id string) (role.RoleResponse, error) {
	var role role.RoleResponse
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		return role, err
	}

	fmt.Println(role)

	return role, nil
}

func (r *RepositoryImpl) GetByRole(c echo.Context, db *gorm.DB, roleId string) ([]role.MemberRoleResponse, error) {
	var role []role.MemberRoleResponse
	if err := db.Where("role_id = ?", roleId).Joins("User").Find(&role).Error; err != nil {
		return role, err
	}

	fmt.Println(role)

	return role, nil
}

func (r *RepositoryImpl) Create(c echo.Context, db *gorm.DB, data models.MRole) (int64, error) {
	var id int64
	if err := db.Create(&data).Error; err != nil {
		return id, err
	}

	id = *data.ID

	return id, nil
}

func (r *RepositoryImpl) Update(c echo.Context, db *gorm.DB, data models.MRole) (int64, error) {
	var existingData models.MRole

	if err := db.First(&existingData, data.ID).Error; err != nil {
		return 0, err
	}

	result := db.Model(&existingData).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected

	return rowsAffected, nil
}

func (r *RepositoryImpl) Delete(c echo.Context, db *gorm.DB, id string) (int64, error) {
	var existingData models.MRole

	if err := db.First(&existingData, id).Error; err != nil {
		return 0, err
	}

	result := db.Delete(&existingData)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected

	return rowsAffected, nil
}
