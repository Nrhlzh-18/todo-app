package controller

import (
	"net/http"

	"github.com/Nrhlzh-18/todo-app/app/tasks/service"
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

	return res.SuccessResponse(c, http.StatusOK, "success get all data tasks", result)
}
