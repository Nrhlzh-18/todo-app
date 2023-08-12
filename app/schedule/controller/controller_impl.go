package controller

import (
	"net/http"

	"github.com/Nrhlzh-18/todo-app/app/schedule"
	"github.com/Nrhlzh-18/todo-app/app/schedule/service"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/labstack/echo/v4"
)

type ControllerImpl struct {
	Service service.Service
}

func NewController(service service.Service) Controller {
	return &ControllerImpl{
		Service: service,
	}
}

func (co *ControllerImpl) GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get all schedule", result)
}

func (co *ControllerImpl) GetById(c echo.Context) error {
	id := c.Param("id")

	result, err := co.Service.GetById(c, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get id schedule", result)
}

func (co *ControllerImpl) Create(c echo.Context) error {
	var data schedule.ScheduleRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	err = co.Service.Create(c, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusCreated, "success menambahkan schedule", nil)
}

func (co *ControllerImpl) Update(c echo.Context) error {
	var data schedule.ScheduleRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	id := c.Param("id")

	result, err := co.Service.Update(c, id, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success memperbarui schedule", result)
}

func (co *ControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	rowsAffected, err := co.Service.Delete(c, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	var response res.ResponseRowsAffected
	response.RowsAffected = uint64(rowsAffected)

	return res.SuccessResponse(c, http.StatusOK, "success menghapus schedule", response)
}
