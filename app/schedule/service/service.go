package service

import (
	"github.com/Nrhlzh-18/todo-app/app/schedule"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]schedule.ScheduleResponse, error)
	GetById(c echo.Context, id string) (models.MSchedule, error)
	GetByDate(c echo.Context) ([]models.MSchedule, error)
	Create(c echo.Context, request schedule.ScheduleRequest) error
	Update(c echo.Context, id string, request schedule.ScheduleRequest) (int64, error)
	Delete(c echo.Context, id string) (int64, error)
}
