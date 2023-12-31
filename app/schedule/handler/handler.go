package handler

import (
	"github.com/Nrhlzh-18/todo-app/app/schedule/controller"
	"github.com/Nrhlzh-18/todo-app/app/schedule/repository"
	"github.com/Nrhlzh-18/todo-app/app/schedule/service"
	repo_user "github.com/Nrhlzh-18/todo-app/app/user/repository"
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
	repo_user := repo_user.NewRepository()
	srv := service.NewService(repo, repo_user, db, validate)
	ctrl := controller.NewController(srv)
	return &handler{
		Controller: ctrl,
	}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Controller.GetAll)
	g.GET("/:id", h.Controller.GetById)
	g.POST("", h.Controller.Create)
	g.PUT("/:id", h.Controller.Update)
	g.DELETE("/:id", h.Controller.Delete)
}
