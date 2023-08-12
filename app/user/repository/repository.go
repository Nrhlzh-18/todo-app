package repository

import (
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]user.UserResponse, error)
	GetById(c echo.Context, db *gorm.DB, id string) (user.UserResponse, error)
	GetByUsernamePass(c echo.Context, db *gorm.DB, user user.UserLogin) (models.MUser, error)
	Create(c echo.Context, db *gorm.DB, data models.MUser) (models.MUser, error)
	Update(c echo.Context, db *gorm.DB, data models.MUser) (models.MUser, error)
	Delete(c echo.Context, db *gorm.DB, id string) error
}
