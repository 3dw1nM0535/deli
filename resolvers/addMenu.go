package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) AddMenu(ctx context.Context, input models1.MenuInput) (*models.Menu, error) {
	// validate input for nullness
	if input.MenuHeadline == "" {
		return &models.Menu{}, errors.New("menu cannot miss a catchy headline")
	}
	var id = utils.ParseUUID(input.RestaurantID) // parse id to uuid
	if id.String() == "" {
		return &models.Menu{}, errors.New("restaurant id for the menu should be provided")
	}

	// check that the restaurant is registered with us
	restaurant := &models.Restaurant{}
	r.ORM.DB.Where("id = ?", id).First(&restaurant)
	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("cannot find restaurant with id '%s'", id.String())
		return &models.Menu{}, err
	}

	// proceed to save the data if clean
	menu := &models.Menu{
		Headline:     input.MenuHeadline,
		RestaurantID: utils.ParseUUID(input.RestaurantID),
	}
	r.ORM.DB.Save(&menu)
	return menu, nil
}
