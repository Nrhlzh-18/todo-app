package models

import "time"

type MTasks struct {
	ID          int64     `json:"id" gorm:"primary_key;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;null"`
	DueDate     time.Time `json:"due_date" gorm:"type:date;"`
	Priority    string    `json:"priority" gorm:"type:varchar(225);default:Low"`
	Status      string    `json:"status" gorm:"type:varchar(225)"`
	CreatedBy   int64     `json:"created_by" gorm:"index;not null"`
	UpdatedBy   int64     `json:"updated_by" gorm:"index;null"`
	Created_at  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	UserCreated MUser `json:"user_created" gorm:"foreignkey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserUpdated MUser `json:"user_updated" gorm:"foreignkey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TasksTag struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	IDTasks    int64     `json:"id_tasks" gorm:"index"`
	IDUser     int64     `json:"id_user" gorm:"foreignKey;null"`
	IDProject  int64     `json:"id_project" gorm:"foreignKey;null"`
	IDTeam     int64     `json:"id_team" gorm:"foreignKey;null"`
	Created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
