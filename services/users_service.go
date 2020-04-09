package services

import (
	"github.com/Hantara/bookstore_user-api/domain/users"
)

//CreateUser = Business logic pembuatan user
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
