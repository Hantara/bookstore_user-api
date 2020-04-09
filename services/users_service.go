package services

import (
	"github.com/Hantara/bookstore_user-api/domain/users"
	"github.com/Hantara/bookstore_user-api/utils/errors"
)

//GetUser = pengambilan user sesuai dengan list
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

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
