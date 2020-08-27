package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
)

func (r *queryResolver) GetHarvestingFarms(ctx context.Context) ([]*models.Farm, error) {
	farms := []*models.Farm{}
	r.ORM.DB.Where("season = ?", "Harvesting").Find(&farms)
	return farms, nil
}
