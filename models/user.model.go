package models

import "time"

type MUser struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(225);not null"`
	Email        string    `json:"email" gorm:"unique;type:varchar(225);not null;"`
	Username     string    `json:"username" gorm:"unique;type:varchar(225);not null"`
	Password     string    `json:"password" gorm:"type:varchar(225);not null"`
	PasswordHash string    `gorm:"unique;type:varchar(225);not null;size:255"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

type MUserRole struct {
	ID     int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser int64  `json:"id_user" gorm:"index"`
	Name   string `json:"name" gorm:"type:varchar(225);not null"`

	User MUser `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type MUserProject struct {
	ID     int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser int64  `json:"id_user" gorm:"index"`
	Name   string `json:"name" gorm:"type:varchar(225);not null"`

	User MUser `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
