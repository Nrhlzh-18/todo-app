package repository

import (
	"github.com/Nrhlzh-18/todo-app/app/comments"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, db *gorm.DB) ([]comments.CommentResponse, error)
	GetById(c echo.Context, db *gorm.DB, id string) (comments.CommentResponse, error)
	GetByProject(c echo.Context, db *gorm.DB, projectID string) ([]comments.CommentResponse, error)
	GetByUser(c echo.Context, db *gorm.DB, userID string) ([]comments.CommentResponse, error)
	Create(c echo.Context, db *gorm.DB, data models.MComments) (comments.CommentResponse, error)
	Update(c echo.Context, db *gorm.DB, data models.MComments) (comments.CommentResponse, error)
	Delete(c echo.Context, db *gorm.DB, id string) error
}
