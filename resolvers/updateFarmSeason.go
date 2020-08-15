package resolvers

import (
	"context"
	"errors"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateFarmSeason(ctx context.Context, input models.SeasonUpdateInput) (*models1.Farm, error) {
	farm := &models1.Farm{}
	r.ORM.DB.Where("id = ?", input.Token).First(&farm)
	if r.ORM.DB.NewRecord(farm) {
		notFound := errors.New("not found")
		return nil, notFound
	}
	r.ORM.DB.Model(&farm).Update("season", input.Season)
	return farm, nil
}
