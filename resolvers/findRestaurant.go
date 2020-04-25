package resolvers

import (
	"context"
	"errors"
	"fmt"

	models1 "github.com/3dw1nM0535/Byte/db/models"
)

// FindRestaurant : find restaurant partner using its id
func (r *queryResolver) FindRestaurant(ctx context.Context, id string) (*models1.Restaurant, error) {
	if id == "" {
		err := errors.New("id field cannot be empty")
		return &models1.Restaurant{}, err
	}
	var restaurant = &models1.Restaurant{}
	r.ORM.DB.Where("id = ?", id).First(&restaurant)
	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("restaurant with id '%s' cannot be found", id)
		return &models1.Restaurant{}, err
	}
	return restaurant, nil
}
