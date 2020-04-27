package resolvers

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) UploadMc(ctx context.Context, input models1.UploadDoc) (*models1.File, error) {
	ctx = context.Background()

	// validate rider is registered
	rider := &models.Rider{}
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.ID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("no rider with id '%s' is registered with Byte", input.ID)
		return &models1.File{}, err
	}

	_, attr, err := utils.Upload(ctx, input.File.File, riderIDBucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	// Parse object URL
	objURL := utils.ObjectURL(attr)
	medicalCert := &models.MDC{
		Media:     objURL,
		Content:   attr.ContentType,
		Size:      int(attr.Size),
		CreatedAt: attr.Created,
		UpdatedAt: attr.Updated,
		RiderID:   utils.ParseUUID(input.ID),
	}
	r.ORM.DB.Save(&medicalCert)
	return &models1.File{
		ID:        medicalCert.ID.String(),
		Media:     medicalCert.Media,
		Content:   medicalCert.Content,
		Size:      medicalCert.Size,
		CreatedAt: &medicalCert.CreatedAt,
		UpdatedAt: &medicalCert.UpdatedAt,
	}, nil
}
