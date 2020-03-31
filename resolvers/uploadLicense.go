package resolvers

import (
	"context"

	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
	"github.com/99designs/gqlgen/graphql"
	"github.com/joho/godotenv"
)

var credPath, projectID, bucketName string

func init() {
	godotenv.Load()
	credPath = utils.MustGetEnv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID = utils.MustGetEnv("GOOGLE_PROJECT_ID")
	bucketName = utils.MustGetEnv("GOOGLE_BUCKET_NAME")
}

func (r *mutationResolver) UploadLicense(ctx context.Context, file graphql.Upload) (*models1.File, error) {
	ctx = context.Background()
	_, attr, err := utils.Upload(ctx, file.File, bucketName, credPath, projectID, file.Filename)
	if err != nil {
		return &models1.File{}, err
	}
	return &models1.File{
		Name:      attr.MediaLink,
		ConteType: attr.ContentType,
	}, nil
}
