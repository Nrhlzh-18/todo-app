package schedule

import "github.com/Nrhlzh-18/todo-app/util"

type ScheduleRequest struct {
	Name        *string       `json:"name" validate:"required"`
	Description *string       `json:"description" validate:"required"`
	Date        util.JsonDate `json:"date"`
	StartTime   util.JsonTime `json:"start_time" validate:"required"`
	EndTime     util.JsonTime `json:"end_time" validate:"required"`
	Location    *string       `json:"location" validate:"required"`
}

type ScheduleResponse struct {
	ID          *int64            `json:"id"`
	Name        *string           `json:"name"`
	Description *string           `json:"description"`
	Date        util.JsonDate     `json:"date"`
	StartTime   util.JsonTime     `json:"start_time"`
	EndTime     util.JsonTime     `json:"end_time"`
	Location    *string           `json:"location"`
	CreatedAt   util.JsonDateTime `json:"created_at"`
	UpdatedAt   util.JsonDateTime `json:"updated_at"`
}

func (*ScheduleRequest) TableName() string {
	return "m_schedule"
}

func (*ScheduleResponse) TableName() string {
	return "m_schedule"
}
