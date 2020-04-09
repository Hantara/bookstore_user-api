package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication menjalankan routing dan maps url
func StartApplication() {
	mapUrls()
	router.Run(":8080")

}
