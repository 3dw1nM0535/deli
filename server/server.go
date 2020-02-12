package server

import (
	"github.com/3dw1nM0535/deli/db"
	handler "github.com/3dw1nM0535/deli/handlers"
	"github.com/gin-gonic/gin"
)

// Run : spin server
func Run(orm *db.DB) {
	r := gin.Default()
	r.POST("/query", handler.GraphQL(orm))
	r.GET("/", handler.Playground())
	r.Run()
}
