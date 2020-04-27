package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"

	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) AddDisplayPics(ctx context.Context, input models1.UploadDocs) ([]*models1.File, error) {
	// validate restaurant id
	id := utils.ParseUUID(input.ID)
	if id.String() == "" {
		return []*models1.File{}, errors.New("restaurant id cannot be empty")
	}
	// check if restaurant exists
	restaurant := &models.Restaurant{}
	r.ORM.DB.First(&restaurant, "id = ?", id)
	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("restaurant with id '%s' not found", input.ID)
		return []*models1.File{}, err
	}
	// Upload to google cloud storage
	files := []*models1.File{}
	for i := range input.Files {
		_, attr, err := utils.Upload(ctx, input.Files[i].File, restaurantDPBucket, credPath, projectID, input.Files[i].Filename)
		if err != nil {
			return []*models1.File{}, err
		}
		objURL := utils.ObjectURL(attr)
		file := &models.DisplayPic{
			Media:        objURL,
			Content:      attr.ContentType,
			Size:         int(attr.Size),
			CreatedAt:    attr.Created,
			UpdatedAt:    attr.Updated,
			RestaurantID: id,
		}
		r.ORM.DB.Save(&file)
		files = append(files, &models1.File{
			ID:        file.ID.String(),
			Media:     file.Media,
			Content:   file.Content,
			Size:      file.Size,
			CreatedAt: &file.CreatedAt,
			UpdatedAt: &file.UpdatedAt,
		})
	}
	return files, nil
}
