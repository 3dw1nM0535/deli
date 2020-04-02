package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

// check if menu exists
func menuExists(id string) bool {
	menu := models.Menu{}
	uid := "00000000-0000-0000-0000-000000000000"
	validID := utils.ParseUUID(id)
	if validID.String() == uid {
		return false
	}
	db, _ := db.Factory()
	db.DB.Where("id = ?", id).First(&menu)
	if menu.ID.String() == uid {
		return false
	}
	return true
}

func mapItemsToDish(items []*models1.DishInput) ([]*models.Dish, error) {
	ctx := context.Background()
	dishes := []*models.Dish{}
	// validate input for null
	if len(items) == 0 {
		return []*models.Dish{}, errors.New("dishes cannot be empty")
	}

	for index := range items {
		if items[index].Title == "" {
			return []*models.Dish{}, errors.New("dish title cannot be empty")
		}
		if items[index].Description == "" {
			return []*models.Dish{}, errors.New("dish description cannot be empty")
		}
		if items[index].Image.Filename == "" {
			return []*models.Dish{}, errors.New("you must provide dish image")
		}
		if fmt.Sprintf("%.2f", float64(items[index].Price)) == "0.00" {
			return []*models.Dish{}, errors.New("dish price must be known to customers")
		}
		if menuExists(items[index].MenuID) == false {
			return []*models.Dish{}, errors.New("dish must belong to a menu. provide a valid menu id")
		}

		file := items[index].Image.File
		fileName := items[index].Image.Filename
		_, attr, err := utils.Upload(ctx, file, dishesBucketName, credPath, projectID, fileName)
		if err != nil {
			return []*models.Dish{}, err
		}

		d := &models.Dish{
			Title:       items[index].Title,
			Description: items[index].Description,
			Price:       items[index].Price,
			Image:       attr.MediaLink,
			AddOns:      items[index].AddOns,
			MenuID:      utils.ParseUUID(items[index].MenuID),
		}
		dishes = append(dishes, d)
	}
	return dishes, nil
}

func (r *mutationResolver) AddDish(ctx context.Context, input []*models1.DishInput) ([]*models.Dish, error) {
	dishes, err := mapItemsToDish(input)
	if err != nil {
		return []*models.Dish{}, err
	}

	for i := range dishes {
		r.ORM.DB.Save(&dishes[i])
	}
	return dishes, nil
}
