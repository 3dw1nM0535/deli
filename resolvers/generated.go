package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db"
	models1 "github.com/3dw1nM0535/deli/db/models"
	graph "github.com/3dw1nM0535/deli/graph/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	ORM *db.ORM
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Restaurant() graph.RestaurantResolver {
	return &restaurantResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type restaurantResolver struct{ *Resolver }

func (r *restaurantResolver) ID(ctx context.Context, obj *models1.Restaurant) (*string, error) {
	id := obj.ID.String()
	return &id, nil
}
