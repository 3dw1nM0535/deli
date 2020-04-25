package resolvers

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) UploadDp(ctx context.Context, input models1.UploadDoc) (*models1.File, error) {
	ctx = context.Background()

	// validate if rider is available
	rider := &models.Rider{}
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.ID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("cannot find rider with id '%s' registered with Byte", input.ID)
		return &models1.File{}, err
	}

	_, attr, err := utils.Upload(ctx, input.File.File, riderDPBucket, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	dp := &models.DisplayPicture{
		Media:     attr.MediaLink,
		Size:      int(attr.Size),
		Content:   attr.ContentType,
		CreatedAt: attr.Created,
		UpdatedAt: attr.Updated,
		RiderID:   utils.ParseUUID(input.ID),
	}

	r.ORM.DB.Save(&dp)
	return &models1.File{
		ID:        dp.ID.String(),
		Media:     dp.Media,
		Content:   dp.Content,
		Size:      dp.Size,
		CreatedAt: &dp.CreatedAt,
		UpdatedAt: &dp.UpdatedAt,
	}, nil
}
