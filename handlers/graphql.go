package handler

import (
	"github.com/3dw1nM0535/Byte/db"
	graph "github.com/3dw1nM0535/Byte/graph/generated"
	resolver "github.com/3dw1nM0535/Byte/resolvers"
	"github.com/99designs/gqlgen/handler"
	option "github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphQL : websocket handler
func GraphQL(orm *db.ORM, options []option.Option) gin.HandlerFunc {
	// Pass Configs
	cfg := graph.Config{
		Resolvers: &resolver.Resolver{
			ORM: orm,
		},
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
