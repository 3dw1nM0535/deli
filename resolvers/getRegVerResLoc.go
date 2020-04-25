package resolvers

import (
	"context"

	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *queryResolver) FindRegVerResLoc(ctx context.Context) ([]*models1.City, error) {
	var cities []*models1.City

	// find registered and verified restaurant location
	query := `SELECT city FROM addresses INNER JOIN restaurant_addresses ON restaurant_addresses.address_id = addresses.id INNER JOIN restaurants ON restaurant_addresses.restaurant_id = restaurants.id WHERE restaurants.verified = true;`
	r.ORM.DB.Raw(query).Scan(&cities)

	return cities, nil
}
