package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
	"strconv"
)

func (r *mutationResolver) AddFarm(ctx context.Context, input models1.FarmInput) (*models.Farm, error) {
	id, _ := strconv.Atoi(string(input.ID))
	newFarm := &models.Farm{
		ID:        id,
		Size:      input.Size,
		Soil:      input.Soil,
		ImageHash: input.ImageHash,
	}
	r.ORM.DB.Save(&newFarm)
	return newFarm, nil
}
