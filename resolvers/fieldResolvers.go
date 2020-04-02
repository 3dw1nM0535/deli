package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
)

func (r *restaurantResolver) ID(ctx context.Context, obj *models.Restaurant) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *dishResolver) ID(ctx context.Context, obj *models.Dish) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *licenseResolver) ID(ctx context.Context, obj *models.License) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *dishResolver) AddOns(ctx context.Context, obj *models.Dish) ([]string, error) {
	var addOns []string
	addOns = obj.AddOns
	return addOns, nil
}

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	id := obj.ID.String()
	return id, nil
}

// Dishes : find dishes belonging to a menu
func (r *menuResolver) Dishes(ctx context.Context, obj *models.Menu) ([]*models.Dish, error) {
	dishes := []*models.Dish{}
	menu := obj
	r.ORM.DB.First(&menu, "id = ?", obj.ID)
	r.ORM.DB.Model(&menu).Related(&dishes)
	return dishes, nil
}

// Addresses : find addresses belonging to a restaurant
func (r *restaurantResolver) Addresses(ctx context.Context, obj *models.Restaurant) ([]*models.Address, error) {
	addresses := []*models.Address{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&addresses, "Address")
	return addresses, nil
}

// Restaurants : find restaurants belonging to an address
func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
	restaurants := []*models.Restaurant{}
	address := obj
	r.ORM.DB.First(&address, "id = ?", obj.ID)
	r.ORM.DB.Model(&address).Related(&restaurants, "Restaurants")
	return restaurants, nil
}

// License : find license belonging to a restaurant
func (r *restaurantResolver) License(ctx context.Context, obj *models.Restaurant) (*models.License, error) {
	license := &models.License{}
	r.ORM.DB.Model(&obj).Related(&license)
	return license, nil
}

// Menu : find menu belonging to a restaurant
func (r *restaurantResolver) Menu(ctx context.Context, obj *models.Restaurant) ([]*models.Menu, error) {
	menu := []*models.Menu{}
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	r.ORM.DB.Model(&restaurant).Related(&menu)
	return menu, nil
}
