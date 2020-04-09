package services

import (
	"strings"

	"github.com/Hantara/bookstore_user-api/domain/users"
	"github.com/Hantara/bookstore_user-api/utils/errors"
)

//CreateUser = Business logic pembuatan konsumen
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, errors.NewBadRequestError("Invalid Email Address")
	}
	return &user, nil
}
