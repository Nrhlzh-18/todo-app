package taskstag

import (
	"github.com/Nrhlzh-18/todo-app/app/tasks"
)

type TasksTagProject struct {
	ID     int64               `gorm:"primaryKey"`
	IDTask int64               `json:"-" gorm:"column:id_tasks"`
	Tasks  tasks.TasksResponse `gorm:"foreignKey:id_tasks;references:ID"`
}

func (TasksTagProject) TableName() string {
	return "tasks_tag"
}

type TasksTagTeam struct {
	ID     int64               `gorm:"primaryKey"`
	IDTask int64               `json:"-" gorm:"column:id_tasks"`
	Tasks  tasks.TasksResponse `gorm:"foreignKey:id_tasks;references:ID"`
}

func (TasksTagTeam) TableName() string {
	return "tasks_tag"
}

type TasksTagUser struct {
	ID     int64               `gorm:"primaryKey"`
	IDTask int64               `json:"-"`
	Tasks  tasks.TasksResponse `gorm:"foreignKey:IDTasks;references:ID"`
}

func (TasksTagUser) TableName() string {
	return "tasks_tag"
}

