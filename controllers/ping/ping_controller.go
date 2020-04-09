package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping adalah controller defaults
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
