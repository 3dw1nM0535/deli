package resolvers

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func (r *mutationResolver) UploadGcc(ctx context.Context, input models1.UploadGcc) (*models1.File, error) {
	ctx = context.Background()

	// validate rider is registered
	rider := &models.Rider{}
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.RiderID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("no rider with id '%s' is registered with Deli", input.RiderID)
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
		RiderID:   utils.ParseUUID(input.RiderID),
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
