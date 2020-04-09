package errors

import "net/http"

//RestErr untuk template rest error
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError mengembalikan template
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
