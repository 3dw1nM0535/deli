package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
)

func (r *restaurantResolver) ID(ctx context.Context, obj *models.Restaurant) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *licenseResolver) ID(ctx context.Context, obj *models.License) (string, error) {
	id := obj.ID.String()
	return id, nil
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
