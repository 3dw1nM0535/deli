package resolvers

import (
	"context"

	models1 "github.com/3dw1nM0535/deli/db/models"
)

func (r *queryResolver) FindRestaurants(ctx context.Context) ([]*models1.Restaurant, error) {
	var restaurants = []*models1.Restaurant{}
	r.ORM.DB.Find(&restaurants)
	return restaurants, nil
}
