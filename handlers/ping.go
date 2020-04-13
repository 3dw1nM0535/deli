package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping : simple keep-alive ping handler
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// Hook : Mpesa hook handler
func Hook() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		c.JSON(http.StatusOK, gin.H{
			"message": string(body),
		})
	}
}
