package controller

import (
	"net/http"

	"github.com/Nrhlzh-18/todo-app/app/auth/service"
	"github.com/Nrhlzh-18/todo-app/app/user"
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

func (co *ControllerImpl) CheckLogin(c echo.Context) error {
	var data user.UserLogin

	if err := c.Bind(&data); err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	err := co.Service.CheckLogin(c, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "berhasil login", nil)
}
