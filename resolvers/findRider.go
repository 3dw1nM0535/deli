package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/utils"
)

func (r *queryResolver) FindRider(ctx context.Context, id string) (*models.Rider, error) {
	rider := &models.Rider{}

	// validate input
	if id == "" {
		return &models.Rider{}, errors.New("rider id cannot be empty")
	}

	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(id))

	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("rider with id '%s' cannot be found", id)
		return &models.Rider{}, err
	}
	return rider, nil
}
