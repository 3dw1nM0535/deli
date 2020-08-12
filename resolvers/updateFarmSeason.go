package resolvers

import (
	"context"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateFarmSeason(ctx context.Context, input models.SeasonUpdateInput) (*models1.Farm, error) {
	farm := &models1.Farm{}
	r.ORM.DB.Where("id = ?", input.Token).First(&farm)
	r.ORM.DB.Model(&farm).Update("season", input.Season)
	return farm, nil
}
