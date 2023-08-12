package controller

import (
	"net/http"

	"github.com/Nrhlzh-18/todo-app/app/tasks"
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

	return res.SuccessResponse(c, http.StatusOK, "success get all tasks", result)
}

func (co *ControllerImpl) GetById(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get by id tasks", result)
}

func (co *ControllerImpl) GetByProject(c echo.Context) error {
	projectID := c.Param("id")
	result, err := co.Service.GetByProject(c, projectID)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get by project tasks", result)
}

func (co *ControllerImpl) GetByTeam(c echo.Context) error {
	teamID := c.Param("id")
	result, err := co.Service.GetByTeam(c, teamID)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get by user tasks", result)
}

func (co *ControllerImpl) GetByUser(c echo.Context) error {
	userID := c.Param("id")
	result, err := co.Service.GetByUser(c, userID)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success get by user tasks", result)
}

func (co *ControllerImpl) Create(c echo.Context) error {
	var data tasks.TaskRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	result, err := co.Service.Create(c, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusCreated, "success menambahkan tasks", result)
}

func (co *ControllerImpl) Update(c echo.Context) error {
	var data tasks.TaskRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	id := c.Param("id")

	result, err := co.Service.Update(c, id, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success memperbarui tasks", result)
}

func (co *ControllerImpl) UpdatePriority(c echo.Context) error {
	id := c.Param("id")
	priority := c.Param("priority")

	err := co.Service.UpdatePriority(c, priority, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success memperbarui tasks", nil)
}

func (co *ControllerImpl) UpdateStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.Param("status")

	err := co.Service.UpdateStatus(c, status, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success memperbarui tasks", nil)
}

func (co *ControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	err := co.Service.Delete(c, id)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusOK, "success menghapus tasks", nil)
}

func (co *ControllerImpl) CreateTags(c echo.Context) error {
	var data tasks.TaksTagRequest
	err := c.Bind(&data)
	if err != nil {
		return res.ErrorResponse(c, res.BuildError(res.ErrServerError, err))
	}

	result, err := co.Service.CreateTags(c, data)
	if err != nil {
		return res.ErrorResponse(c, err)
	}

	return res.SuccessResponse(c, http.StatusCreated, "success menambahkan tasks", result)
}
