package resolvers

import (
	"context"
	"errors"

	models1 "github.com/3dw1nM0535/deli/db/models"
)

func (r *queryResolver) FindRestaurant(ctx context.Context, id *string) (*models1.Restaurant, error) {
	var restaurant = &models1.Restaurant{}
	r.ORM.DB.Where("id = ?", id).First(&restaurant)
	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return &models1.Restaurant{}, errors.New("cannot find the request restaurant")
	}
	return restaurant, nil
}
