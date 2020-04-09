package app

import (
	"github.com/Hantara/bookstore_user-api/Controllers/ping"
	"github.com/Hantara/bookstore_user-api/Controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
