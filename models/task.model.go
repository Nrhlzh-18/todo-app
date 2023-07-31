package models

import "time"

type MTasks struct {
	ID          int64     `json:"id" gorm:"primary_key;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;null"`
	DueDate     time.Time `json:"due_date" gorm:"type:datetime;"`
	Priority    string    `json:"priority" gorm:"type:char"`
	Status      string    `json:"status" gorm:"type:char"`
	Created_at  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

type TasksTag struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	IDRole     int64     `json:"id_role" gorm:"index;null"`
	IDUser     int64     `json:"id_user" gorm:"index;null"`
	Created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	Role MUserRole `json:"role" gorm:"foreignkey:IDRole;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User MUser     `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
