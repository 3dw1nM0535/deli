package resolvers

import (
	"context"
	"errors"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateFarmHarvestSupply(ctx context.Context, input models1.HarvestUpdateInput) (*models.Season, error) {
	season := &models.Season{}
	r.ORM.DB.Where("season_number = ? AND token = ?", input.SeasonNumber, input.Token).Find(&season)
	if season.ID.String() == specialUUID && r.ORM.DB.NewRecord(season) {
		notFound := errors.New("not found")
		return nil, notFound
	}
	season.HarvestYield = input.NewSupply
	r.ORM.DB.Save(&season)
	return season, nil
}
