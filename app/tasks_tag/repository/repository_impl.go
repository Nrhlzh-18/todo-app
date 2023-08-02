package repository

import (
	taskstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetByProject(c echo.Context, db *gorm.DB, projectID string) ([]taskstag.TasksTagProject, error) {
	var tasks []taskstag.TasksTagProject
	if err := db.Where("id_project", projectID).Joins("Tasks").Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *RepositoryImpl) GetByTeam(c echo.Context, db *gorm.DB, teamID string) ([]taskstag.TasksTagTeam, error) {
	var tasks []taskstag.TasksTagTeam
	if err := db.Where("id_team", teamID).Joins("Tasks").Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *RepositoryImpl) GetByUser(c echo.Context, db *gorm.DB, UserID string) ([]taskstag.TasksTagUser, error) {
	var tasks []taskstag.TasksTagUser
	if err := db.Where("id_user", UserID).Joins("Tasks").Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}
