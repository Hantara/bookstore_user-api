package services

import (
	"os/user"

	"github.com/Hantara/bookstore_user-api/domain/users"
)

//CreateUser = Business logic pembuatan user
func CreateUser(user users.User) (*user.User, error) {
	return &user, nil
}
