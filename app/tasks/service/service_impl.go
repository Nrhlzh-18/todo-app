package service

import (
	"time"

	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/Nrhlzh-18/todo-app/app/tasks/repository"
	taskstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag"
	"github.com/Nrhlzh-18/todo-app/helpers"

	repo_tag "github.com/Nrhlzh-18/todo-app/app/tasks_tag/repository"

	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	Repository repository.Repository
	Repo_tag   repo_tag.Repository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewService(
	repository repository.Repository,
	repo_taskstag repo_tag.Repository,
	db *gorm.DB,
	validate *validator.Validate,
) Service {
	return &ServiceImpl{
		Repository: repository,
		Repo_tag:   repo_taskstag,
		DB:         db,
		Validate:   validate,
	}
}

func (s *ServiceImpl) GetAll(c echo.Context) ([]tasks.TasksResponse, error) {
	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetById(c echo.Context, id string) (tasks.TasksResponse, error) {
	result, err := s.Repository.GetById(c, s.DB, id)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetByTeam(c echo.Context, teamID string) ([]taskstag.TasksTagTeam, error) {
	result, err := s.Repo_tag.GetByTeam(c, s.DB, teamID)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetByProject(c echo.Context, projectID string) ([]taskstag.TasksTagProject, error) {
	result, err := s.Repo_tag.GetByProject(c, s.DB, projectID)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetByUser(c echo.Context, userID string) ([]taskstag.TasksTagUser, error) {
	result, err := s.Repo_tag.GetByUser(c, s.DB, userID)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) Create(c echo.Context, request tasks.TaskRequest) (tasks.TasksCreateResponse, error) {
	var response tasks.TasksCreateResponse
	err := s.Validate.Struct(request)
	if err != nil {
		return response, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	payload := models.MTasks{
		Name:        request.Name,
		Description: request.Description,
		DueDate:     time.Time(request.DueDate),
		Priority:    request.Priority,
		Status:      request.Status,
		CreatedBy:   1,
		UpdatedBy:   1,
	}

	data, err := s.Repository.Create(c, s.DB, payload)
	if err != nil {
		return response, res.BuildError(res.ErrServerError, err)
	}

	response.LastInsertID = uint64(data.ID)

	return response, nil
}

func (s *ServiceImpl) Update(c echo.Context, id string, request tasks.TaskRequest) (tasks.TasksResponse, error) {
	var response tasks.TasksResponse
	err := s.Validate.Struct(request)
	if err != nil {
		return response, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	indInt, err := helpers.StringToInt64(id)
	if err != nil {
		return response, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	payload := models.MTasks{
		ID:          indInt,
		Name:        request.Name,
		Description: request.Description,
		DueDate:     time.Time(request.DueDate),
		Priority:    request.Priority,
		Status:      request.Status,
		UpdatedBy:   1,
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
