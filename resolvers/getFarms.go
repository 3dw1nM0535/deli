package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/db/models"
)

func (r *queryResolver) GetFarms(ctx context.Context) ([]*models1.Farm, error) {
	farms := []*models.Farm{}
	returnFarms := []*models1.Farm{}
	r.ORM.DB.Find(&farms)
	for i := 0; i < len(farms); i++ {
		f := &models1.Farm{
			ID:        farms[i].ID,
			Size:      farms[i].Size,
			Soil:      farms[i].Soil,
			ImageHash: farms[i].ImageHash,
			Season:    farms[i].Season,
			Owner:     farms[i].Owner,
		}
		returnFarms = append(returnFarms, f)
	}
	return returnFarms, nil
}
