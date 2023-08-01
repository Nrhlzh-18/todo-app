package service

import (
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/app/user/repository"
	"github.com/Nrhlzh-18/todo-app/helpers"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/Nrhlzh-18/todo-app/models"
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

func (s *ServiceImpl) GetAll(c echo.Context) ([]user.UserResponse, error) {
	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetById(c echo.Context, id string) (user.UserResponse, error) {
	result, err := s.Repository.GetById(c, s.DB, id)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) Create(c echo.Context, request user.UserRequest) (user.UserCreateResponse, error) {
	var response user.UserCreateResponse
	err := s.Validate.Struct(request)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	hashpass, err := helpers.HashPassword(request.Password)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}
	payload := models.MUser{
		Name:         request.Name,
		Email:        request.Email,
		Username:     request.Username,
		Password:     request.Password,
		PasswordHash: hashpass,
	}

	_, err = s.Repository.Create(c, s.DB, payload)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	return response, nil
}

func (s *ServiceImpl) Update(c echo.Context, id string, request user.UserRequest) (user.UserResponse, error) {
	var response user.UserResponse
	err := s.Validate.Struct(request)
	if err != nil {
		return response, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	indInt, err := helpers.StringToInt64(id)
	if err != nil {
		return response, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	hashpass, err := helpers.HashPassword(request.Password)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	payload := models.MUser{
		ID:           indInt,
		Name:         request.Name,
		Email:        request.Email,
		Username:     request.Username,
		Password:     request.Password,
		PasswordHash: hashpass,
	}

	_, err = s.Repository.Update(c, s.DB, payload)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	return response, nil
}

func (s *ServiceImpl) Delete(c echo.Context, id string) error {
	err := s.Repository.Delete(c, s.DB, id)
	if err != nil {
		return res.BuildError(res.ErrServerError, err)
	}

	return nil
}
