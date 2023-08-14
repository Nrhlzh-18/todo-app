package models

import (
	"time"

	"github.com/Nrhlzh-18/todo-app/util"
)

type MSchedule struct {
	ID          *int64         `json:"ID" gorm:"primary_key;autoIncrement"`
	Name        *string        `json:"name" gorm:"type:varchar(225)"`
	Description *string        `json:"description" gorm:"type:text;null"`
	Date        *time.Time     `json:"date" gorm:"type:date;not null"`
	StartTime   *util.JsonTime `json:"start_time" gorm:"type:time;null"`
	EndTime     *util.JsonTime `json:"end_time" gorm:"type:time;null"`
	Location    *string        `json:"location" gorm:"type:varchar(225)"`
	CreatedAt   *time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
