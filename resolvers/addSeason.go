package resolvers

import (
	"context"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *seasonResolver) ID(ctx context.Context, obj *models1.Season) (string, error) {
	id := obj.ID.String()
	return id, nil
}

func (r *mutationResolver) AddSeason(ctx context.Context, input models.SeasonInput) (*models1.Season, error) {
	seasonData := &models1.Season{
		Token:         input.Token,
		Crop:          *input.Crop,
		Fertilizer:    *input.Fertilizer,
		Seed:          *input.Seed,
		ExpectedYield: *input.ExpectedYield,
		SeedSupplier:  *input.SeedSupplier,
		HarvestYield:  *input.HarvestYield,
		HarvestPrice:  *input.HarvestPrice,
	}
	r.ORM.DB.Save(&seasonData)
	return seasonData, nil
}
