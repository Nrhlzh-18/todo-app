package tasks

import (
	"github.com/Nrhlzh-18/todo-app/app/project"
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/helpers"
)

type TaskRequest struct {
	IDProject   int64            `json:"id_project"`
	IDUser      int64            `json:"id_user"`
	Name        string           `json:"name" validate:"required"`
	Description string           `json:"description"`
	DueDate     helpers.JsonDate `json:"due_date" validate:"required"`
	Priority    string           `json:"priority"`
	Status      string           `json:"status" validate:"required"`
	CreatedBy   int64            `json:"created_by"`
	UpdatedBy   int64            `json:"updated_by"`
}

type TasksResponse struct {
	ID          uint64           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	DueDate     helpers.JsonDate `json:"due_date"`
	Priority    string           `json:"priority"`
	Status      string           `json:"status"`
	CreatedAt   helpers.JsonTime `json:"created_at"`
	UpdatedAt   helpers.JsonTime `json:"updated_at"`
}

func (TasksResponse) TableName() string {
	return "m_tasks"
}

type TasksCreateResponse struct {
	LastInsertID uint64 `json:"last_insert_id"`
}

type TasksTag struct {
	ID        int64                   `gorm:"primaryKey"`
	IDTask    int64                   `json:"-"`
	Tasks     user.UserResponse       `gorm:"ForeignKey:IDTasks;references:ID"`
	IDProject int64                   `json:"-"`
	Project   project.ProjectResponse `gorm:"ForeignKey:IDProject;references:ID"`
}
