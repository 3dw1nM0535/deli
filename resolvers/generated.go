package resolvers

import (
	"context"

	graph "github.com/3dw1nM0535/deli/graph/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver : root resolver
type Resolver struct{}

// Query : root query
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Hello : return 'Hello, World!'
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}
