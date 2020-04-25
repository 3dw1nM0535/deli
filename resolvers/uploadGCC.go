package resolvers

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) UploadGcc(ctx context.Context, input models1.UploadDoc) (*models1.File, error) {
	ctx = context.Background()

	// validate rider is registered
	rider := &models.Rider{}
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.ID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("no rider with id '%s' is registered with Byte", input.ID)
		return &models1.File{}, err
	}

	_, attr, err := utils.Upload(ctx, input.File.File, riderGCCBucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	goodConductCert := &models.GCC{
		Media:     attr.MediaLink,
		Content:   attr.ContentType,
		Size:      int(attr.Size),
		CreatedAt: attr.Created,
		UpdatedAt: attr.Updated,
		RiderID:   utils.ParseUUID(input.ID),
	}
	r.ORM.DB.Save(&goodConductCert)
	return &models1.File{
		ID:        goodConductCert.ID.String(),
		Media:     goodConductCert.Media,
		Content:   goodConductCert.Content,
		Size:      goodConductCert.Size,
		CreatedAt: &goodConductCert.CreatedAt,
		UpdatedAt: &goodConductCert.UpdatedAt,
	}, nil
}
