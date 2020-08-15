package resolvers

import (
	"context"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

const (
	specialUUID = "00000000-0000-0000-0000-000000000000"
)

func (r *mutationResolver) UpdateFarmPreparations(ctx context.Context, input models.PreparationInput) (*models1.Season, error) {
	season := &models1.Season{}
	r.ORM.DB.Where("season_number = ? AND token = ?", input.SeasonNumber, input.Token).First(&season)
	if season.ID.String() == specialUUID && r.ORM.DB.NewRecord(season) {
		season.Crop = input.Crop
		season.SeasonNumber = input.SeasonNumber
		season.Token = input.Token
		season.Fertilizer = input.Fertilizer
		r.ORM.DB.Create(&season)
		return season, nil
	}
	r.ORM.DB.Model(&season).Updates(&models1.Season{
		Crop:       input.Crop,
		Fertilizer: input.Fertilizer,
	})
	return season, nil
}
