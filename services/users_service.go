package services

import (
	"github.com/Hantara/bookstore_user-api/domain/users"
	"github.com/Hantara/bookstore_user-api/utils/errors"
)

//CreateUser = Business logic pembuatan konsumen
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
