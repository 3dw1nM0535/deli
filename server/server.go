package server

import (
	"log"

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
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/query", handler.GraphQL())
	r.GET("/graphql", handler.Playground())
	r.GET("/", handler.Ping())
	return r
}

// Run : spins the server
func Run() {
	r := SetupRouter()

	log.Fatal(r.Run(host + ":" + port))
}
