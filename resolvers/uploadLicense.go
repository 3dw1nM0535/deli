package resolvers

import (
	"context"
	"errors"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
	"github.com/joho/godotenv"
)

var credPath,
	projectID,
	licenseBucketName,
	dishesBucketName,
	geocodingKey,
	token,
	riderIDBucketName,
	riderGCCBucketName,
	riderMDCBucketName,
	riderDPBucket string

func init() {
	godotenv.Load()
	credPath = utils.MustGetEnv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID = utils.MustGetEnv("GOOGLE_PROJECT_ID")
	licenseBucketName = utils.MustGetEnv("LICENSE_BUCKET")
	dishesBucketName = utils.MustGetEnv("DISHES_BUCKET")
	geocodingKey = utils.MustGetEnv("GOOGLE_GEOCODING_KEY")
	riderGCCBucketName = utils.MustGetEnv("RIDER_GCC_BUCKET")
	riderIDBucketName = utils.MustGetEnv("RIDER_IDD_BUCKET")
	riderMDCBucketName = utils.MustGetEnv("RIDER_MDC_BUCKET")
	riderDPBucket = utils.MustGetEnv("RIDER_DP_BUCKET")

	orderCreatedChannel = map[string]map[string]chan *models.Order{}
	// token = utils.GetToken()
}

func (r *mutationResolver) UploadLicense(ctx context.Context, input models1.UploadLicense) (*models1.File, error) {
	// validate id
	if input.RestaurantID == "" {
		return &models1.File{}, errors.New("restaurant id cannot be empty")
	}

	var id = utils.ParseUUID(input.RestaurantID)

	// check that restaurant is already in our database
	// before attaching its license
	var restaurant = &models.Restaurant{}
	r.ORM.DB.First(&restaurant, "id = ?", id)

	if restaurant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return &models1.File{}, errors.New("no such restaurant to add its license")
	}

	// check if restaurant already has license attached to it
	var file = &models.License{}
	r.ORM.DB.First(&file, "restaurant_id = ?", id)
	if file.RestaurantID == id {
		return &models1.File{}, errors.New("single business permit cannot have multiple licenses")
	}

	// upload to google cloud storage and return object attributes
	ctx = context.Background()
	_, attr, err := utils.Upload(ctx, input.File.File, licenseBucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	var license = &models.License{
		Media:        attr.MediaLink,
		Content:      attr.ContentType,
		Size:         int(attr.Size),
		CreatedAt:    attr.Created,
		UpdatedAt:    attr.Updated,
		RestaurantID: id,
		Restaurant:   restaurant,
	}
	r.ORM.DB.Save(&license)

	return &models1.File{
		ID:        license.ID.String(),
		Media:     license.Media,
		Content:   license.Content,
		Size:      license.Size,
		CreatedAt: &license.CreatedAt,
		UpdatedAt: &license.UpdatedAt,
	}, nil
}
