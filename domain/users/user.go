package users

import (
	"strings"

	"github.com/Hantara/bookstore_user-api/utils/errors"
)

//User plain object
type User struct {
	Id          int16  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate - validasi input
func Validate(user *User) *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
