package service

import (
	"time"

	"github.com/Nrhlzh-18/todo-app/app/schedule"
	"github.com/Nrhlzh-18/todo-app/app/schedule/repository"
	repo_user "github.com/Nrhlzh-18/todo-app/app/user/repository"
	"github.com/Nrhlzh-18/todo-app/helpers"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	Repository repository.Repository
	RepoUser   repo_user.Repository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewService(
	repository repository.Repository,
	repo_user repo_user.Repository,
	db *gorm.DB,
	validate *validator.Validate,
) Service {
	return &ServiceImpl{
		Repository: repository,
		RepoUser:   repo_user,
		DB:         db,
		Validate:   validate,
	}
}

func (s *ServiceImpl) GetAll(c echo.Context) ([]models.MSchedule, error) {
	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetById(c echo.Context, id string) (models.MSchedule, error) {
	result, err := s.Repository.GetById(c, s.DB, id)
	if err != nil {
		return result, res.BuildError(res.ErrServerError, err)
	}

	return result, nil
}

func (s *ServiceImpl) GetByDate(c echo.Context) ([]models.MSchedule, error) {

	result, err := s.Repository.GetByDate(c, s.DB)
	if err != nil {
		return result, err
	}

	return result, nil
}


func (s *ServiceImpl) Create(c echo.Context, request schedule.ScheduleRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		return res.BuildError(res.ErrServerError, err)
	}

	payload := &models.MSchedule{
		Name:        request.Name,
		Description: request.Description,
		Date:        (*time.Time)(&request.Date),
		StartTime:   (*time.Time)(&request.StartTime),
		EndTime:     (*time.Time)(&request.EndTime),
		Location:    request.Location,
	}

	err = s.Repository.Store(c, s.DB, payload)
	if err != nil {
		return res.BuildError(res.ErrServerError, err)
	}

	return nil
}

func (s *ServiceImpl) Update(c echo.Context, id string, request schedule.ScheduleRequest) (int64, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return 0, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	_, err = helpers.StringToInt64(id)
	if err != nil {
		return 0, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	payload := &models.MSchedule{
		Name:        request.Name,
		Description: request.Description,
		Date:        (*time.Time)(&request.Date),
		StartTime:   (*time.Time)(&request.StartTime),
		EndTime:     (*time.Time)(&request.EndTime),
		Location:    request.Location,
	}

	rowsAffected, err := s.Repository.Update(c, s.DB, payload)
	if err != nil {
		return rowsAffected, res.BuildError(res.ErrServerError, err)
	}

	return rowsAffected, nil
}

func (s *ServiceImpl) Delete(c echo.Context, id string) (int64, error) {
	rowsAffected, err := s.Repository.Delete(c, s.DB, id)
	if err != nil {
		return rowsAffected, res.BuildError(res.ErrServerError, err)
	}

	return rowsAffected, nil
}
