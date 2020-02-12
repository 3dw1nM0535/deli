package resolvers

import (
	"context"

	graph "github.com/3dw1nM0535/deli/graph/generated"
	"github.com/3dw1nM0535/deli/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	delis []*models.Deli
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
	r.delis = append(r.delis, deli)
	return deli, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetDeli(ctx context.Context) ([]*models.Deli, error) {
	return r.delis, nil
}
