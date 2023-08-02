package comments

import (
	"github.com/Nrhlzh-18/todo-app/app/schedule"
	"github.com/Nrhlzh-18/todo-app/app/tasks"
	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/helpers"
)

type CommentRequest struct {
	IDTasks    int64  `json:"id_tasks"`
	IDSchedule int64  `json:"id_schedule"`
	IDUser     int64  `json:"id_user" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
	IDParent   int64  `json:"id_parent"`
}

type CommentResponse struct {
	ID         int64                     `json:"id"`
	IDTasks    tasks.TasksResponse       `json:"id_tasks"`
	IDSchedule schedule.ScheduleResponse `json:"id_schedule"`
	IDUser     user.UserResponse         `json:"id_user"`
	Comment    string                    `json:"comment"`
	IDParent   int64                     `json:"id_parent"`
	CreatedAt  helpers.JsonTime          `json:"created_at"`
	UpdatedAt  helpers.JsonTime          `json:"updated_at"`
}
