package models

type MProject struct {
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(225);not null"`
}

type UserProject struct {
	ID        int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser    int64 `json:"id_user" gorm:"index"`
	IDProject int64 `json:"id_project" gorm:"index"`

	User    MUser    `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project MProject `json:"project" gorm:"foreignkey:IDProject;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
