package resolvers

import (
	"context"
	"time"

	"github.com/3dw1nM0535/deli/db"
	"github.com/3dw1nM0535/deli/db/models"
	graph "github.com/3dw1nM0535/deli/graph/generated"
	models1 "github.com/3dw1nM0535/deli/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *db.DB
}

func (r *Resolver) Deli() graph.DeliResolver {
	return &deliResolver{r}
}
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type deliResolver struct{ *Resolver }

func (r *deliResolver) ID(ctx context.Context, obj *models.Deli) (string, error) {
	return obj.ID.String(), nil
}
func (r *deliResolver) Delicacies(ctx context.Context, obj *models.Deli) ([]string, error) {
	return obj.Delicacies, nil
}
func (r *deliResolver) Reviews(ctx context.Context, obj *models.Deli) ([]string, error) {
	return obj.Reviews, nil
}
func (r *deliResolver) CreatedAt(ctx context.Context, obj *models.Deli) (*time.Time, error) {
	return &obj.CreatedAt, nil
}
func (r *deliResolver) UpdatedAt(ctx context.Context, obj *models.Deli) (*time.Time, error) {
	return &obj.UpdatedAt, nil
}
func (r *deliResolver) DeletedAt(ctx context.Context, obj *models.Deli) (*time.Time, error) {
	return &obj.DeletedAt, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddDeli(ctx context.Context, input models1.AddDeli) (*models.Deli, error) {
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
	var delis []*models.Deli
	r.DB.Raw("SELECT * FROM delis").Scan(&delis)
	return delis, nil
}
