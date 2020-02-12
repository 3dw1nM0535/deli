package resolvers

import (
	"context"

	"github.com/3dw1nM0535/deli/db"
	graph "github.com/3dw1nM0535/deli/graph/generated"
	"github.com/3dw1nM0535/deli/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *db.DB
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddDeli(ctx context.Context, input models.AddDeli) (*models.Deli, error) {
	deli := &models.Deli{
		RestaurantName: input.RestaurantName,
		Telephone:      input.Telephone,
		Delicacies:     input.Delicacies,
	}
	r.DB.Create(&deli)
	return deli, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetDeli(ctx context.Context) ([]*models.Deli, error) {
	delis := []*models.Deli{}
	r.DB.Find(&delis)
	return delis, nil
}
