package resolvers

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
)

var credPath, projectID, bucketName string

func init() {
	godotenv.Load()
	credPath = utils.MustGetEnv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID = utils.MustGetEnv("GOOGLE_PROJECT_ID")
	bucketName = utils.MustGetEnv("LICENSE_BUCKET_NAME")
}

func (r *mutationResolver) UploadLicense(ctx context.Context, input models1.UploadLicense) (*models.License, error) {
	// validate id
	if input.RestaurantID == "" {
		return &models.License{}, errors.New("restaurant id cannot be empty")
	}

	var id = uuid.Must(uuid.FromString(input.RestaurantID))

	// check that restaurant is already in our database
	// before attaching its license
	var restaurant = &models.Restaurant{}
	r.ORM.DB.First(&restaurant, "id = ?", id)

	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return &models.License{}, errors.New("no such restaurant to add its license")
	}

	// check if restaurant already has license attached to it
	var lice = &models.License{}
	r.ORM.DB.First(&lice, "restaurant_id = ?", id)
	if lice.RestaurantID == id {
		return &models.License{}, errors.New("single business permit cannot have multiple licenses")
	}

	// upload to google cloud storage and return object attributes
	ctx = context.Background()
	_, attr, err := utils.Upload(ctx, input.File.File, bucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models.License{}, err
	}

	var license = &models.License{
		Media:        attr.MediaLink,
		Content:      attr.ContentType,
		Size:         attr.Size,
		CreatedAt:    attr.Created,
		UpdatedAt:    attr.Updated,
		RestaurantID: id,
	}
	r.ORM.DB.Save(&license)
	return license, nil
}
