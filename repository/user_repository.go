package repository

import (
	"github.com/hariszaki17/go-api-clean/entity"
)

// UserRepository expose global
type UserRepository interface {
	Insert(user entity.User)
	FindAll() (users []entity.User)
	DeleteAll()
	Encrypt(plainText string) (chiper string)
	Decrypt(password, chiper string)
	ValidatePassword(username, password string)
}