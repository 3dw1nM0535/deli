package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
	graph "github.com/3dw1nM0535/deli/graph/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	ORM *db.ORM
}

func (r *Resolver) Address() graph.AddressResolver {
	return &addressResolver{r}
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

type addressResolver struct{ *Resolver }

func (r *addressResolver) RestaurantID(ctx context.Context, obj *models.Address) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
// 	panic("not implemented")
// }

type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) AddRestaurant(ctx context.Context, input models1.RestaurantInput) (*models.Restaurant, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) RegisterAddress(ctx context.Context, input models1.AddressInput) (*models.Address, error) {
// 	panic("not implemented")
// }

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Hello(ctx context.Context) (string, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindRestaurant(ctx context.Context, id string) (*models.Restaurant, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindRestaurants(ctx context.Context) ([]*models.Restaurant, error) {
// 	panic("not implemented")
// }

type restaurantResolver struct{ *Resolver }

func (r *restaurantResolver) ID(ctx context.Context, obj *models.Restaurant) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// func (r *restaurantResolver) Addresses(ctx context.Context, obj *models.Restaurant) ([]*models.Address, error) {
// 	panic("not implemented")
// }
