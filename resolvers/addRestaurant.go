package resolvers

import (
	"context"
	"errors"
	"fmt"

	models1 "github.com/3dw1nM0535/deli/db/models"
	"github.com/3dw1nM0535/deli/models"
)

// ID : return ID as string value
func (r *restaurantResolver) ID(ctx context.Context, obj *models1.Restaurant) (*string, error) {
	id := obj.ID.String()
	return &id, nil
}

// AddRestaurant : create and save restaurant to the database
func (r *mutationResolver) AddRestaurant(ctx context.Context, input models.RestaurantInput) (*models1.Restaurant, error) {
	// validate required input information
	if input.RestaurantName == "" {
		err := errors.New("restaurant name cannot be empty")
		return &models1.Restaurant{}, err
	}
	if input.About == "" {
		err := errors.New("restaurant brief description cannot be empty")
		return &models1.Restaurant{}, err
	}
	if input.Telephone == "" {
		err := errors.New("restaurant contact information cannot be empty")
		return &models1.Restaurant{}, err
	}
	// check restaurant duplication
	var restaurant models1.Restaurant
	r.ORM.DB.Where("restaurant_name = ?", input.RestaurantName).First(&restaurant)
	if restaurant.RestaurantName == input.RestaurantName {
		err := fmt.Errorf("restaurant with business name '%s' already exists", input.RestaurantName)
		return &models1.Restaurant{}, err
	}
	// create restaurant if not registered earlier
	var newRestaurant = &models1.Restaurant{
		RestaurantName: input.RestaurantName,
		About:          input.About,
		Telephone:      input.Telephone,
	}
	r.ORM.DB.Create(&newRestaurant)
	return newRestaurant, nil
}
