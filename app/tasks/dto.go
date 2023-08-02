package tasks

import (
	"github.com/Nrhlzh-18/todo-app/helpers"
)

type TaskRequest struct {
	Name        string           `json:"name" validate:"required"`
	Description string           `json:"description"`
	DueDate     helpers.JsonDate `json:"due_date" validate:"required"`
	Priority    string           `json:"priority"`
	Status      string           `json:"status" validate:"required"`
	CreatedBy   int64            `json:"created_by"`
	UpdatedBy   int64            `json:"updated_by"`
}

type TasksResponse struct {
	ID          uint64               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	DueDate     helpers.JsonDate     `json:"due_date"`
	Priority    string               `json:"priority"`
	Status      string               `json:"status"`
}

func (TasksResponse) TableName() string {
	return "m_tasks"
}

type TasksCreateResponse struct {
	LastInsertID uint64 `json:"last_insert_id"`
}
