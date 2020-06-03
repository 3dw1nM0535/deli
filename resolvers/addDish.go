package resolvers

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/Byte/db"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

// check if menu exists
func menuExists(id string) bool {
	menu := models.Menu{}
	uid := "00000000-0000-0000-0000-000000000000"
	validID := utils.ParseUUID(id)
	if validID.String() == uid {
		return false
	}
	db, err := db.Factory()
	// check for any error
	if err != nil {
		return false
	}
	db.DB.Where("id = ?", validID).First(&menu)
	if menu.ID.String() == uid {
		return false
	}
	return true
}

// AddDish : add dish controller
func (r *mutationResolver) AddDish(ctx context.Context, input models1.DishInput) (*models.Dish, error) {
	// validate data
	if input.MenuID == "" {
		return &models.Dish{}, errors.New("dish menu id cannot be empty")
	}
	if input.Title == "" {
		return &models.Dish{}, errors.New("dish title cannot be empty")
	}
	if input.Price == 0 {
		return &models.Dish{}, errors.New("dish cannot sell at 0")
	}
	// check if dish menu provided exists
	if !menuExists(input.MenuID) {
		return &models.Dish{}, errors.New("dish menu does not exist")
	}

	// Upload dish image
	_, attr, err := utils.Upload(ctx, input.Image.File, dishesBucketName, credPath, projectID, input.Image.Filename)
	if err != nil {
		return &models.Dish{}, err
	}

	// proceed to creating dish
	objURL := utils.ObjectURL(attr)
	dish := &models.Dish{
		Title:       input.Title,
		Description: input.Description,
		Image:       objURL,
		Price:       input.Price,
		MenuID:      utils.ParseUUID(input.MenuID),
	}
	r.ORM.DB.Save(&dish)
	return dish, nil
}
