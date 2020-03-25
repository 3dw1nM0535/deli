package server

import (
	"log"

	"github.com/3dw1nM0535/deli/db"
	handler "github.com/3dw1nM0535/deli/handlers"
	"github.com/3dw1nM0535/deli/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var host, port string

func init() {
	godotenv.Load()
	host = utils.MustGetEnv("SERVER_HOST")
	port = utils.MustGetEnv("SERVER_PORT")
}

// SetupRouter : set routing paths
func SetupRouter(orm *db.ORM) *gin.Engine {
	r := gin.Default()
	r.POST("/query", handler.GraphQL(orm))
	r.GET("/graphql", handler.Playground())
	r.GET("/", handler.Ping())
	return r
}

// Run : spins the server
func Run(orm *db.ORM) {
	r := SetupRouter(orm)

	log.Fatal(r.Run(host + ":" + port))
}
