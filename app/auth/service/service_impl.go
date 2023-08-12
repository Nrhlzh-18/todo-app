package service

import (
	"fmt"
	"time"

	"github.com/Nrhlzh-18/todo-app/app/user"
	"github.com/Nrhlzh-18/todo-app/app/user/repository"
	repo_user "github.com/Nrhlzh-18/todo-app/app/user/repository"
	"github.com/Nrhlzh-18/todo-app/helpers"
	res "github.com/Nrhlzh-18/todo-app/helpers/response"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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

func (s *ServiceImpl) CheckLogin(c echo.Context, user user.UserLogin) (string, error) {

	var password string
	password = *user.Password
	result, err := s.Repository.GetByUsernamePass(c, s.DB, user)
	if err != nil {
		return "", res.BuildError(res.ErrServerError, err)
	}

	var pwd string
	pwd = *result.PasswordHash

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match.")
		return "", &res.ErrNotFound
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = result.ID
	claims["name"] = result.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", &res.ErrUnauthorized
	}

	return t, nil
}
