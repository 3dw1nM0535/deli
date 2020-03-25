package resolvers

import (
	"context"

	models1 "github.com/3dw1nM0535/deli/db/models"
	"github.com/3dw1nM0535/deli/models"
)

func (r *mutationResolver) AddRestaurant(ctx context.Context, input models.RestaurantInput) (*models1.Restaurant, error) {
	var newRestaurant = &models1.Restaurant{
		RestaurantName: input.RestaurantName,
		About:          input.About,
		Telephone:      input.Telephone,
	}
	r.ORM.DB.Create(&newRestaurant)
	return newRestaurant, nil
}
