package handler

import (
	"github.com/3dw1nM0535/deli/db"
	graph "github.com/3dw1nM0535/deli/graph/generated"
	resolver "github.com/3dw1nM0535/deli/resolvers"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphQL : /graphql handler
func GraphQL(orm *db.DB) gin.HandlerFunc {
	// Pass Configs
	cfg := graph.Config{
		Resolvers: &resolver.Resolver{
			DB: orm,
		},
	}

	h := handler.GraphQL(graph.NewExecutableSchema(cfg))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground : return graphql playground
func Playground() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
