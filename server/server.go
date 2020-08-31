package server

import (
	"log"
	"net/http"
	"time"

	"github.com/3dw1nM0535/Byte/db"
	handler "github.com/3dw1nM0535/Byte/handlers"
	"github.com/3dw1nM0535/Byte/utils"
	option "github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var port string
var options []option.Option
var c gin.HandlerFunc

func init() {
	godotenv.Load()
	port = utils.MustGetEnv("PORT")

	c = cors.Default()

	options = []option.Option{
		option.WebsocketKeepAliveDuration(10 * time.Second),
		option.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}),
	}
}

// SetupRouter : set routing paths
func SetupRouter(orm *db.ORM) *gin.Engine {
	r := gin.Default()
	r.Use(c)
	r.POST("/query", handler.GraphQL(orm, options))
	r.GET("/graphql", handler.Playground())
	r.GET("/", handler.Ping())
	r.GET("/ws", handler.GraphQL(orm, options))
	return r
}

// Run : spins the server
func Run(orm *db.ORM) {
	r := SetupRouter(orm)

	s := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	s.SetKeepAlivesEnabled(true)
	log.Fatal(s.ListenAndServe())
}
