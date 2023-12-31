package tasks

import (
	"github.com/Nrhlzh-18/todo-app/util"
)

type TaskRequest struct {
	Name        *string       `json:"name" validate:"required"`
	Description *string       `json:"description"`
	DueDate     util.JsonDate `json:"due_date" validate:"required"`
	Priority    *string       `json:"priority"`
	Status      *string       `json:"status" validate:"required"`
	CreatedBy   *string       `json:"created_by"`
	UpdatedBy   *string       `json:"updated_by"`
}

type TasksResponse struct {
	ID          *int64        `json:"id"`
	Name        *string       `json:"name"`
	Description *string       `json:"description"`
	DueDate     util.JsonDate `json:"due_date"`
	Priority    *string       `json:"priority"`
	Status      *string       `json:"status"`
}

func (TasksResponse) TableName() string {
	return "m_tasks"
}

type TasksCreateResponse struct {
	LastInsertID *uint64 `json:"last_insert_id"`
}

type TaksTagRequest struct {
	IDTasks   *int64 `json:"id_tasks" validate:"required"`
	IDUser    *int64 `json:"id_user" validate:"required"`
	IDProject *int64 `json:"id_project"`
	IDTeam    *int64 `json:"id_team"`
}
