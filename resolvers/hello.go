package resolvers

import "context"

// Hello : return hello world
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}
