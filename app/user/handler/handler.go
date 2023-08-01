package handler

import (
	"github.com/Nrhlzh-18/todo-app/app/user/controller"
	"github.com/Nrhlzh-18/todo-app/app/user/repository"
	"github.com/Nrhlzh-18/todo-app/app/user/service"
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
	g.GET("/:id", h.Controller.GetById)
	g.POST("/create", h.Controller.Create)
	g.PUT("/update/:id", h.Controller.Update)
	g.DELETE("/delete/:id", h.Controller.Delete)
}
