package handler

import (
	"github.com/Nrhlzh-18/todo-app/app/auth/controller"
	"github.com/Nrhlzh-18/todo-app/app/auth/service"
	"github.com/Nrhlzh-18/todo-app/app/user/repository"
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
	r_user := repo_user.NewRepository()
	repo := repository.NewRepository()
	srv := service.NewService(repo, db, validate, r_user)
	ctrl := controller.NewController(srv)
	return &handler{
		Controller: ctrl,
	}
}

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Controller.CheckLogin)
}
