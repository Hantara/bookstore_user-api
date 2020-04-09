package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Hantara/bookstore_user-api/domain/users"
	"github.com/Hantara/bookstore_user-api/services"
	"github.com/gin-gonic/gin"
)

//GetUser controller untuk melakukan reply
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Under development")
}

//CreateUser controller untuk melakukan reply
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: handle error
		fmt.Println("byteskosong")
		fmt.Println(err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: handle json error
		fmt.Println("json")
		fmt.Println(err.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO : handle create user error
		fmt.Println("createuser")
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}

//SearchUser controller untuk melakukan reply
/* func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Under development")
} */
