package models

type MTeam struct {
	ID   *int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name *string `json:"name" gorm:"type:varchar(225);not null"`
}

type UserTeam struct {
	ID     *int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser *int64 `json:"id_user" gorm:"index"`
	IDTeam *int64 `json:"id_team" gorm:"index"`

	User *MUser `json:"user" gorm:"foreignkey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Team *MTeam `json:"team" gorm:"foreignkey:IDTeam;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
