package resolvers

import (
	"context"
)

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, Welcome to Octopus!", nil
}
