package resolvers

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) AddDishAddOn(ctx context.Context, input models1.DishAddOnInput) (*models.DishAddOn, error) {
	// validate input
	if input.DishID == "" {
		return &models.DishAddOn{}, errors.New("dish id cannot be empty")
	}
	if input.Name == "" {
		return &models.DishAddOn{}, errors.New("dish name cannot be empty")
	}
	if input.Price == 0 {
		return &models.DishAddOn{}, errors.New("dish cannot sell at 0")
	}
	// Check if dish exists
	dish := &models.Dish{}
	r.ORM.DB.First(&dish, "id = ?", utils.ParseUUID(input.DishID))
	if dish.ID.String() == "" {
		return &models.DishAddOn{}, errors.New("dish not found")
	}

	// Proceed to creating dish add on
	addOn := &models.DishAddOn{
		Name:   input.Name,
		Price:  input.Price,
		DishID: utils.ParseUUID(input.DishID),
	}
	r.ORM.DB.Save(&addOn)

	return addOn, nil
}
