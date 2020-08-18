package resolvers

import (
	"context"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *queryResolver) GetSeasons(ctx context.Context, input models.SeasonsQueryInput) ([]*models1.Season, error) {
	seasons := []*models1.Season{}
	r.ORM.DB.Where("token = ?", input.Token).Find(&seasons)
	return seasons, nil
}
