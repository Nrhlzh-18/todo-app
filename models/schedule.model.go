package models

import "time"

type MSchedule struct {
	ID          int64     `json:"ID" gorm:"primary_key;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(225)"`
	Description string    `json:"description" gorm:"type:text;null"`
	Date        time.Time `json:"date" gorm:"type:date;not null"`
	StartTime   time.Time `json:"start_time" gorm:"type:time;not null"`
	EndTime     time.Time `json:"end_time" gorm:"type:time;not null"`
	Location    string    `json:"location" gorm:"type:varchar(225)"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
