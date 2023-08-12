package schedule

import (
	"github.com/Nrhlzh-18/todo-app/helpers"
)

type ScheduleRequest struct {
	Name        *string           `json:"name" validate:"required"`
	Description *string           `json:"description" validate:"required"`
	Date        helpers.JsonDate `json:"date" validate:"required"`
	StartTime   helpers.JsonTime `json:"start_time"`
	EndTime     helpers.JsonTime `json:"end_time"`
	Location    *string           `json:"location" validate:"required"`
}

type ScheduleResponse struct {
	ID          *int64                `json:"id"`
	Name        *string               `json:"name"`
	Description *string               `json:"description"`
	Date        helpers.JsonDate     `json:"date"`
	StartTime   helpers.JsonTime     `json:"start_time"`
	EndTime     helpers.JsonTime     `json:"end_time"`
	Location    *string               `json:"location"`
	CreatedAt   helpers.JsonDateTime `json:"created_at"`
	UpdatedAt   helpers.JsonDateTime `json:"updated_at"`
}

func (*ScheduleRequest) TableName() string {
	return "m_schedule"
}