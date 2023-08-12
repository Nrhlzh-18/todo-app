package user

type UserRequest struct {
	ID       *int64  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserResponse struct {
	ID       *int64  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserResponseName struct {
	ID       *int64  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}

type UserLogin struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserCreateResponse struct {
	LastInsertID *string `json:"last_insert_id"`
}

func (UserResponse) TableName() string {
	return "m_user"
}
