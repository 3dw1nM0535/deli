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
	if season.ID.String() == specialUUID {
		newSeason := &models1.Season{
			SeasonNumber:  input.SeasonNumber,
			Token:         input.Token,
			Crop:          input.Crop,
			Fertilizer:    input.Fertilizer,
			Seed:          "",
			ExpectedYield: "",
			SeedSupplier:  "",
			HarvestYield:  "",
			HarvestPrice:  "",
		}
		r.ORM.DB.Save(&newSeason)
		return newSeason, nil
	}
	season.Crop = input.Crop
	season.Fertilizer = input.Fertilizer
	r.ORM.DB.Save(&season)
	return season, nil
}
