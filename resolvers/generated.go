package resolvers

import (
	"github.com/3dw1nM0535/deli/db"
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

// func (r *mutationResolver) AddRestaurant(ctx context.Context, input models.RestaurantInput) (*models1.Restaurant, error) {
// 	panic("not implemented")
// }

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Hello(ctx context.Context) (string, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindRestaurant(ctx context.Context, id *string) (*models1.Restaurant, error) {
// 	panic("not implemented")
// }

type restaurantResolver struct{ *Resolver }

// func (r *restaurantResolver) ID(ctx context.Context, obj *models1.Restaurant) (*string, error) {
// 	panic("not implemented")
// }
