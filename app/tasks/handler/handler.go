package handler

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks/controller"
	"github.com/Nrhlzh-18/todo-app/app/tasks/repository"
	"github.com/Nrhlzh-18/todo-app/app/tasks/service"
	tagstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag/repository"
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
	repo_tag := tagstag.NewRepository()
	srv := service.NewService(repo, repo_tag, db, validate)
	ctrl := controller.NewController(srv)
	return &handler{
		Controller: ctrl,
	}
}

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Controller.GetAll)
	g.GET("/:id", h.Controller.GetById)
	g.GET("/user/:id", h.Controller.GetByUser)
	g.GET("/team/:id", h.Controller.GetByTeam)
	g.GET("/project/:id", h.Controller.GetByProject)
	g.POST("/", h.Controller.Create)
	g.POST("/tags", h.Controller.Create)
	g.PUT("/:id", h.Controller.Update)
	g.DELETE("/:id", h.Controller.Delete)
}
