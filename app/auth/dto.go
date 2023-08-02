package auth

import "github.com/Nrhlzh-18/todo-app/helpers"

type RegisterRequest struct {
	Access_token string
	Id_token     string
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Provider  string `json:"provider"`
	Photo     string `json:"photo"`
	Verified  bool   `json:"verified"`
	CreatedAt helpers.JsonTime
	UpdatedAt helpers.JsonTime
}
