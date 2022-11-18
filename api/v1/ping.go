package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestPing(c *gin.Context) {
	// resp := Response{}
	// resp.Message = "pong"
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}