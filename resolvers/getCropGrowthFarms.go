package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
)

func (r *queryResolver) GetCropGrowthFarms(ctx context.Context) ([]*models.Farm, error) {
	farms := []*models.Farm{}
	r.ORM.DB.Where("season = ?", "Crop Growth").Find(&farms)
	return farms, nil
}
