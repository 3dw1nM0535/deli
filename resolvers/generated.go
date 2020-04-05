package resolvers

import (
	"github.com/3dw1nM0535/deli/db"
	graph "github.com/3dw1nM0535/deli/graph/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	ORM *db.ORM
}

func (r *Resolver) Address() graph.AddressResolver {
	return &addressResolver{r}
}
func (r *Resolver) Dish() graph.DishResolver {
	return &dishResolver{r}
}
func (r *Resolver) DishOrder() graph.DishOrderResolver {
	return &dishOrderResolver{r}
}
func (r *Resolver) License() graph.LicenseResolver {
	return &licenseResolver{r}
}
func (r *Resolver) Menu() graph.MenuResolver {
	return &menuResolver{r}
}
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Order() graph.OrderResolver {
	return &orderResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Restaurant() graph.RestaurantResolver {
	return &restaurantResolver{r}
}

type addressResolver struct{ *Resolver }

// func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
// 	panic("not implemented")
// }

type dishResolver struct{ *Resolver }

// func (r *dishResolver) ID(ctx context.Context, obj *models.Dish) (string, error) {
// 	panic("not implemented")
// }
// func (r *dishResolver) AddOns(ctx context.Context, obj *models.Dish) ([]string, error) {
// 	panic("not implemented")
// }

type dishOrderResolver struct{ *Resolver }

// func (r *dishOrderResolver) ID(ctx context.Context, obj *models.DishOrder) (string, error) {
// 	panic("not implemented")
// }
// func (r *dishOrderResolver) AddOns(ctx context.Context, obj *models.DishOrder) ([]string, error) {
// 	panic("not implemented")
// }

type licenseResolver struct{ *Resolver }

// func (r *licenseResolver) ID(ctx context.Context, obj *models.License) (string, error) {
// 	panic("not implemented")
// }

type menuResolver struct{ *Resolver }

// func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
// 	panic("not implemented")
// }
// func (r *menuResolver) Dishes(ctx context.Context, obj *models.Menu) ([]*models.Dish, error) {
// 	panic("not implemented")
// }

type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) AddRestaurant(ctx context.Context, input models1.RestaurantInput) (*models.Restaurant, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) RegisterAddress(ctx context.Context, input models1.AddressInput) (*models.Address, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) UploadLicense(ctx context.Context, input models1.UploadLicense) (*models.License, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) AddMenu(ctx context.Context, input models1.MenuInput) (*models.Menu, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) AddDish(ctx context.Context, input []*models1.DishInput) ([]*models.Dish, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) MakeOrder(ctx context.Context, input models1.OrderInput) (*models.Order, error) {
// 	panic("not implemented")
// }

type orderResolver struct{ *Resolver }

// func (r *orderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
// 	panic("not implemented")
// }
// func (r *orderResolver) Notes(ctx context.Context, obj *models.Order) ([]*models.DishOrder, error) {
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

// func (r *restaurantResolver) ID(ctx context.Context, obj *models.Restaurant) (string, error) {
// 	panic("not implemented")
// }
// func (r *restaurantResolver) Addresses(ctx context.Context, obj *models.Restaurant) ([]*models.Address, error) {
// 	panic("not implemented")
// }
// func (r *restaurantResolver) License(ctx context.Context, obj *models.Restaurant) (*models.License, error) {
// 	panic("not implemented")
// }
// func (r *restaurantResolver) Menu(ctx context.Context, obj *models.Restaurant) ([]*models.Menu, error) {
// 	panic("not implemented")
// }
// func (r *restaurantResolver) Orders(ctx context.Context, obj *models.Restaurant) ([]*models.Order, error) {
// 	panic("not implemented")
// }
