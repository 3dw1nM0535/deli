package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func (r *mutationResolver) UploadID(ctx context.Context, input models1.UploadID) (*models1.File, error) {
	ctx = context.Background()

	_, attr, err := utils.Upload(ctx, input.File.File, riderIDBucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	identificationDoc := &models.IDD{
		Media:     attr.MediaLink,
		Content:   attr.ContentType,
		Size:      int(attr.Size),
		CreatedAt: attr.Created,
		UpdatedAt: attr.Updated,
		RiderID:   utils.ParseUUID(input.RiderID),
	}
	r.ORM.DB.Save(&identificationDoc)
	return &models1.File{
		ID:        identificationDoc.ID.String(),
		Media:     identificationDoc.Media,
		Content:   identificationDoc.Content,
		Size:      identificationDoc.Size,
		CreatedAt: &identificationDoc.CreatedAt,
		UpdatedAt: &identificationDoc.UpdatedAt,
	}, nil
}
