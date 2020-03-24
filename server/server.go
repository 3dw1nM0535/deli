package server

import (
	handler "github.com/3dw1nM0535/deli/handlers"
	"github.com/gin-gonic/gin"
)

// Run : spin server
func Run() {
	r := gin.Default()
	r.POST("/query", handler.GraphQL())
	r.GET("/", handler.Playground())
	r.Run()
}
