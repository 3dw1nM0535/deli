package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping : keep/alive handler
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
