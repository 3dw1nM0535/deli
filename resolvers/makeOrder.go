package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func mapItemsToNotes(items []*models1.DishNote) []*models.DishOrder {
	dishInput := []*models.DishOrder{}

	for index := range items {
		dI := &models.DishOrder{
			DishID:      utils.ParseUUID(items[index].DishID),
			Title:       items[index].Title,
			Description: items[index].Description,
			AddOns:      items[index].AddOns,
			OrderID:     utils.ParseUUID(models.ID),
		}
		dishInput = append(dishInput, dI)
	}
	return dishInput
}

func (r *mutationResolver) MakeOrder(ctx context.Context, input models1.OrderInput) (*models.Order, error) {
	order := &models.Order{
		RestaurantNotes: input.RestaurantNotes,
		RestaurantID:    utils.ParseUUID(input.RestaurantID),
	}
	r.ORM.DB.Save(&order)

	notes := mapItemsToNotes(input.OrderNotes)
	for i := range notes {
		r.ORM.DB.Save(notes[i])
	}
	return order, nil
}
