package repository

import (
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]models.MTasks, error)
	// Create(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error)
	// Update(c echo.Context, db *gorm.DB, data models.MTasks) (models.MTasks, error)
	// Delete(c echo.Context, db *gorm.DB, data models.MTasks) error
}
