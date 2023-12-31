package controller

import (
	"net/http"

	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/app/user/service"
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

	return res.SuccessResponse(c, http.StatusOK, "success get all users", result)
}

func (co *ControllerImpl) GetById(c echo.Context) error {
	id := c.Param("id")

	result, err := co.Service.GetById(c, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get id user", result)
}

func (co *ControllerImpl) Create(c echo.Context) error {
	var data user.UserRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	result, err := co.Service.Create(c, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusCreated, "success menambahkan user", result)
}

func (co *ControllerImpl) Update(c echo.Context) error {
	var data user.UserRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	id := c.Param("id")

	result, err := co.Service.Update(c, id, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success memperbarui user", result)
}

func (co *ControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	err := co.Service.Delete(c, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success menghapus user", nil)
}
