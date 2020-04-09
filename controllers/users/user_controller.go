package users

import (
	"net/http"

	"github.com/Hantara/bookstore_user-api/domain/users"
	"github.com/Hantara/bookstore_user-api/services"
	"github.com/Hantara/bookstore_user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//GetUser controller untuk melakukan reply
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Under development")
}

//CreateUser controller untuk melakukan reply
func CreateUser(c *gin.Context) {
	var user users.User
	//TODO : handle json return bad request
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO : handle create user error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//SearchUser controller untuk melakukan reply
/* func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Under development")
} */
