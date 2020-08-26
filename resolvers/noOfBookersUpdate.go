package resolvers

import (
	"context"
	"errors"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateHarvestBookers(ctx context.Context, input models1.HarvestBookersInput) (*models.Season, error) {
	season := &models.Season{}
	r.ORM.DB.Where("token = ? AND season_number = ?", input.Token, input.SeasonNumber).Find(&season)
	if season.ID.String() == specialUUID && r.ORM.DB.NewRecord(season) {
		notFound := errors.New("not found")
		return nil, notFound
	}
	season.NoOfBookers = season.NoOfBookers + input.NoOfBookers
	r.ORM.DB.Save(&season)
	return season, nil
}
