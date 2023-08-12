package repository

import (
	taskstag "github.com/Nrhlzh-18/todo-app/app/tasks_tag"
	"github.com/Nrhlzh-18/todo-app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetById(c echo.Context, db *gorm.DB, id string) (taskstag.TasksTagProject, error) {
	var task taskstag.TasksTagProject
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
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

func (r *RepositoryImpl) Store(c echo.Context, db *gorm.DB, data models.TasksTag) (int64, error) {
	result := db.Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected

	return rowsAffected, nil
}

func (r *RepositoryImpl) Update(c echo.Context, db *gorm.DB, data models.TasksTag) (int64, error) {
	result := db.Model(&data).Updates(&data)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected
	return rowsAffected, nil
}

func (r *RepositoryImpl) Delete(c echo.Context, db *gorm.DB, id string) (int64, error) {
	var task models.TasksTag
	if err := db.First(&task, id).Error; err != nil {
		return 0, err
	}

	result := db.Delete(&task)
	if result.Error != nil {
		return 0, result.Error
	}

	rowsAffected := result.RowsAffected
	return rowsAffected, nil
}
