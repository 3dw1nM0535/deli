package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/3dw1nM0535/deli/db"
	graph "github.com/3dw1nM0535/deli/graph/generated"
)

type Resolver struct {
	ORM *db.ORM
}

// func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
// 	panic("not implemented")
// }

// func (r *dishResolver) ID(ctx context.Context, obj *models.Dish) (string, error) {
// 	panic("not implemented")
// }

// func (r *dishResolver) AddOns(ctx context.Context, obj *models.Dish) ([]string, error) {
// 	panic("not implemented")
// }

// func (r *dishOrderResolver) ID(ctx context.Context, obj *models.DishOrder) (string, error) {
// 	panic("not implemented")
// }

// func (r *dishOrderResolver) AddOns(ctx context.Context, obj *models.DishOrder) ([]string, error) {
// 	panic("not implemented")
// }

// func (r *licenseResolver) ID(ctx context.Context, obj *models.License) (string, error) {
// 	panic("not implemented")
// }

// func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
// 	panic("not implemented")
// }

// func (r *menuResolver) Dishes(ctx context.Context, obj *models.Menu) ([]*models.Dish, error) {
// 	panic("not implemented")
// }

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

// func (r *orderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
// 	panic("not implemented")
// }

// func (r *orderResolver) Notes(ctx context.Context, obj *models.Order) ([]*models.DishOrder, error) {
// 	panic("not implemented")
// }

// func (r *queryResolver) Hello(ctx context.Context) (string, error) {
// 	panic("not implemented")
// }

// func (r *queryResolver) FindRestaurant(ctx context.Context, id string) (*models.Restaurant, error) {
// 	panic("not implemented")
// }

// func (r *queryResolver) FindNearByRestaurants(ctx context.Context, input models1.Cords) ([]*models.Restaurant, error) {
// 	panic("not implemented")
// }

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

// func (r *subscriptionResolver) OrderCreated(ctx context.Context, id string) (<-chan *models.Order, error) {
// 	panic("not implemented")
// }

// Address returns graph.AddressResolver implementation.
func (r *Resolver) Address() graph.AddressResolver { return &addressResolver{r} }

// Dish returns graph.DishResolver implementation.
func (r *Resolver) Dish() graph.DishResolver { return &dishResolver{r} }

// DishOrder returns graph.DishOrderResolver implementation.
func (r *Resolver) DishOrder() graph.DishOrderResolver { return &dishOrderResolver{r} }

// License returns graph.LicenseResolver implementation.
func (r *Resolver) License() graph.LicenseResolver { return &licenseResolver{r} }

// Menu returns graph.MenuResolver implementation.
func (r *Resolver) Menu() graph.MenuResolver { return &menuResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Order returns graph.OrderResolver implementation.
func (r *Resolver) Order() graph.OrderResolver { return &orderResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Restaurant returns graph.RestaurantResolver implementation.
func (r *Resolver) Restaurant() graph.RestaurantResolver { return &restaurantResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type addressResolver struct{ *Resolver }
type dishResolver struct{ *Resolver }
type dishOrderResolver struct{ *Resolver }
type licenseResolver struct{ *Resolver }
type menuResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type restaurantResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
