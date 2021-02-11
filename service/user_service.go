package service

import "github.com/hariszaki17/go-api-clean/model"

// UserService expose global
type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse)
	List() (responses []model.GetUserResponse)
	DeleteAll() (response model.DeleteAllUserResponse)
	validatePassword(username, password string)
	Login(request model.LoginUserRequest) (response model.LoginUserResponse)
}