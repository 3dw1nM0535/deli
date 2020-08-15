package resolvers

import (
	"context"
	"errors"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateFarmPlantings(ctx context.Context, input models.PlantingInput) (*models1.Season, error) {
	season := &models1.Season{}
	r.ORM.DB.Where("season_number = ? AND token = ?", input.SeasonNumber, input.Token).First(&season)
	if season.ID.String() == specialUUID {
		notFound := errors.New("no season found. start from preparations")
		return nil, notFound
	}
	r.ORM.DB.Model(&season).Updates(&models1.Season{
		Seed:          input.SeedUsed,
		SeedSupplier:  input.SeedSupplier,
		ExpectedYield: input.ExpectedYield,
	})
	return season, nil
}
