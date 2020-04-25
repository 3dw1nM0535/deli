package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *mutationResolver) AddDeliveryMeans(ctx context.Context, input models1.DeliveryMeansInput) (*models1.DeliveryMeans, error) {
	rider := &models.Rider{}

	// validate means input
	if input.Means == "" {
		return &models1.DeliveryMeans{}, errors.New("rider delivery means cannot be empty")
	}

	// validate rider exists
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.RiderID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("cannot find rider with id '%s'", input.RiderID)
		return &models1.DeliveryMeans{}, err
	}

	// validate rider cannot have multiple delivery means
	means := &models.DeliveryMeans{}
	r.ORM.DB.First(&means, "rider_id = ?", utils.ParseUUID(input.RiderID))
	if means.RiderID.String() == input.RiderID {
		return &models1.DeliveryMeans{}, errors.New("rider can only have one deliery means at the time")
	}

	newMeans := &models.DeliveryMeans{
		RiderID: utils.ParseUUID(input.RiderID),
		Means:   input.Means,
	}
	r.ORM.DB.Save(&newMeans)
	return &models1.DeliveryMeans{
		Means: newMeans.Means,
	}, nil
}
