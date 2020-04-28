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
	db, _ := db.Factory()
	db.DB.Where("id = ?", validID).First(&menu)
	if menu.ID.String() == uid {
		return false
	}
	return true
}

// func mapItemsToDish(items []*models1.DishInput) ([]*models.Dish, error) {
// 	ctx := context.Background()
// 	dishes := []*models.Dish{}
// 	// validate input for null
// 	if len(items) == 0 {
// 		return []*models.Dish{}, errors.New("dishes cannot be empty")
// 	}

// 	for i := range items {
// 		if items[i].Title == "" {
// 			return []*models.Dish{}, errors.New("dish title cannot be empty")
// 		}
// 		if items[i].Description == "" {
// 			return []*models.Dish{}, errors.New("dish description cannot be empty")
// 		}
// 		if items[i].Image.Filename == "" {
// 			return []*models.Dish{}, errors.New("you must provide dish image")
// 		}
// 		if fmt.Sprintf("%.2f", float64(items[i].Price)) == "0.00" {
// 			return []*models.Dish{}, errors.New("dish price must be known to customers")
// 		}
// 		if menuExists(items[i].MenuID) == false {
// 			return []*models.Dish{}, errors.New("dish must belong to a menu. provide a valid menu id")
// 		}

// 		file := items[i].Image.File
// 		fileName := items[i].Image.Filename
// 		_, attr, err := utils.Upload(ctx, file, dishesBucketName, credPath, projectID, fileName)
// 		if err != nil {
// 			return []*models.Dish{}, err
// 		}

// 		d := &models.Dish{
// 			Title:       items[i].Title,
// 			Description: items[i].Description,
// 			Price:       items[i].Price,
// 			Image:       attr.MediaLink,
// 			MenuID:      utils.ParseUUID(items[i].MenuID),
// 		}
// 		dishes = append(dishes, d)
// 	}
// 	return dishes, nil
// }

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
