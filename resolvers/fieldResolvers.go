package resolvers

import (
	"context"
	"log"

	"github.com/3dw1nM0535/deli/db/models"
)

// Addresses : find addresses belonging to a restaurant
func (r *restaurantResolver) Addresses(ctx context.Context, obj *models.Restaurant) ([]*models.Address, error) {
	var addresses []*models.Address
	restaurant := obj
	r.ORM.DB.First(&restaurant, "id = ?", obj.ID)
	log.Println(restaurant)
	r.ORM.DB.Model(&restaurant).Related(&addresses, "Address")
	return addresses, nil
}

// Restaurants : find restaurants belonging to an address
func (r *addressResolver) Restaurants(ctx context.Context, obj *models.Address) ([]*models.Restaurant, error) {
	var restaurants []*models.Restaurant
	address := obj
	r.ORM.DB.First(&address, "id = ?", obj.ID)
	r.ORM.DB.Model(&address).Related(&restaurants, "Restaurants")
	return restaurants, nil
}
