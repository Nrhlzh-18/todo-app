package service

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/Nrhlzh-18/todo-app/app/tasks/repository"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	Repository repository.Repository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewService(
	repository repository.Repository,
	db *gorm.DB,
	validate *validator.Validate,
) Service {
	return &ServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (s *ServiceImpl) GetAll(c echo.Context) ([]tasks.TasksResponse, error) {
	var response []tasks.TasksResponse
	_, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	return response, nil
}
