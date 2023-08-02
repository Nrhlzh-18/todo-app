package repository

import (
	"github.com/Nrhlzh-18/todo-app/app/role"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]role.RoleResponse, error)
	GetById(c echo.Context, db *gorm.DB, id string) (role.RoleResponse, error)
	GetByRole(c echo.Context, db *gorm.DB, roleId string) ([]role.MemberRoleResponse, error)
	Create(c echo.Context, db *gorm.DB, data models.MRole) (int64, error)
	Update(c echo.Context, db *gorm.DB, data models.MRole) (int64, error)
	Delete(c echo.Context, db *gorm.DB, id string) (int64, error)
}
