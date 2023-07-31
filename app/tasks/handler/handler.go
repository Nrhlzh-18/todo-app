package handler

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks/controller"
	"github.com/Nrhlzh-18/todo-app/app/tasks/repository"
	"github.com/Nrhlzh-18/todo-app/app/tasks/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	Controller controller.Controller
}

func NewHandler(db *gorm.DB) *handler {
	validate := validator.New()
	repo := repository.NewRepository()
	srv := service.NewService(repo, db, validate)
	ctrl := controller.NewController(srv)
	return &handler{
		Controller: ctrl,
	}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Controller.GetAll)
	g.POST("/create", h.Controller.Create)
}
