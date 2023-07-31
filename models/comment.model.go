package models

import "time"

type MComments struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	IDTasks    int64     `json:"id_tasks" gorm:"index"`
	IDSchedule int64     `json:"id_schedule" gorm:"index"`
	IDUser     int64     `json:"id_user" gorm:"index"`
	Comment    string    `json:"comment" gorm:"type:text;not null"`
	IDParent   int64     `json:"id_parent" gorm:"index"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	Tasks    MTasks     `json:"tasks" gorm:"foreignkey:IDTasks;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedule MSchedule  `json:"schedule" gorm:"foreignkey:IDSchedule;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User     MUser      `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Parent   *MComments `json:"parent" gorm:"foreignkey:IDParent;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
