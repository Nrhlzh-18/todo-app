package models

import "time"

type MTasks struct {
	ID          int64     `json:"id" gorm:"primary_key;autoIncrement"`
	IDProject   int64     `json:"id_project" gorm:"index"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;null"`
	DueDate     time.Time `json:"due_date" gorm:"type:date;"`
	Priority    string    `json:"priority" gorm:"type:varchar(225);default:green"`
	Status      string    `json:"status" gorm:"type:varchar(225)"`
	CreatedBy   int64     `json:"created_by" gorm:"index;not null"`
	UpdatedBy   int64     `json:"updated_by" gorm:"index;null"`
	Created_at  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	Project MUserProject `json:"project" gorm:"foreignkey:IDProject;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    MUser        `json:"user" gorm:"foreignkey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
