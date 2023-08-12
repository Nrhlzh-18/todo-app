package role

import "github.com/Nrhlzh-18/todo-app/app/user"

type RoleRequest struct {
	IDUser *int64  `json:"id_user"`
	Name   *string `json:"name" validate:"required"`
}

type RoleResponse struct {
	IDUser *int64  `json:"id_user"`
	Name   *string `json:"name" validate:"required"`
}

type MemberRoleResponse struct {
	ID     *int64                 `json:"id_role"`
	IDUser *int64                 `json:"-" gorm:"field:id_user"`
	User   *user.UserResponseName `json:"user" gorm:"foreignKey:IDUser;references:ID"`
	Name   *string                `json:"name" validate:"required"`
}
