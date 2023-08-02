package project

type ProjectResponse struct {
	ID   int64  `json:"ID" gorm:"primaryKey"`
	Name string `json:"name"`
}
