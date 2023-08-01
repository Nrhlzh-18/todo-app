package models

import (
	"time"

	"github.com/Nrhlzh-18/todo-app/helpers"
)

type MSchedule struct {
	ID          int64              `json:"ID" gorm:"primary_key;autoIncrement"`
	Name        string             `json:"name" gorm:"type:varchar(225)"`
	Description string             `json:"description" gorm:"type:text;null"`
	Date        time.Time          `json:"date" gorm:"type:date;not null"`
	StartTime   helpers.CustomTime `json:"start_time"`
	EndTime     helpers.CustomTime `json:"end_time"`
	Location    string             `json:"location" gorm:"type:varchar(225)"`
	CreatedAt   time.Time          `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time          `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
