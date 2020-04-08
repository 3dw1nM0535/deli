package handler

import (
	"net/http"
	"time"

	"github.com/3dw1nM0535/deli/db"
	graph "github.com/3dw1nM0535/deli/graph/generated"
	resolver "github.com/3dw1nM0535/deli/resolvers"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GraphQL : websocket handler
func GraphQL(orm *db.ORM) gin.HandlerFunc {
	// Pass Configs
	cfg := graph.Config{
		Resolvers: &resolver.Resolver{
			ORM: orm,
		},
	}

	options := []handler.Option{
		handler.WebsocketKeepAliveDuration(40 * time.Second),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}),
	}

	h := handler.GraphQL(graph.NewExecutableSchema(cfg), options...)

	// srv := h.New(graph.NewExecutableSchema(cfg))

	// srv.AddTransport(transport.GET{})
	// srv.AddTransport(transport.POST{})
	// srv.AddTransport(transport.Websocket{
	// 	KeepAlivePingInterval: 10 * time.Second,
	// 	Upgrader: websocket.Upgrader{
	// 		CheckOrigin: func(r *http.Request) bool {
	// 			return true
	// 		},
	// 	},
	// })

	// srv.Use(extension.Introspection{})

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground : return graphql playground
func Playground() gin.HandlerFunc {
	h := handler.Playground("Interactive GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
