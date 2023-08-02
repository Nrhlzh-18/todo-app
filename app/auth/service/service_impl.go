package service

import (
	"fmt"

	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/app/user/repository"
	repo_user "github.com/Nrhlzh-18/todo-app/app/user/repository"
	"github.com/Nrhlzh-18/todo-app/helpers"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	Repository repository.Repository
	DB         *gorm.DB
	Validate   *validator.Validate
	MUser      repo_user.Repository
}

func NewService(
	repository repository.Repository,
	db *gorm.DB,
	validate *validator.Validate,
	r_user repo_user.Repository,
) Service {
	return &ServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
		MUser:      r_user,
	}
}

func (s *ServiceImpl) CheckLogin(c echo.Context, user user.UserLogin) (bool, error) {

	var password string
	password = user.Password
	result, err := s.Repository.GetByUsernamePass(c, s.DB, &user)
	if err != nil {
		return false, res.BuildError(res.ErrServerError, err)
	}

	var pwd string
	pwd = result.PasswordHash

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match.")
		return false, nil
	}

	return true, nil
}
