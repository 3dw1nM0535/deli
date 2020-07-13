package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/3dw1nM0535/Byte/db"
	graph "github.com/3dw1nM0535/Byte/graph/generated"
)

type Resolver struct {
	ORM *db.ORM
}

/*
 *func (r *mutationResolver) AddFarm(ctx context.Context, input models.FarmInput) (*models1.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) Hello(ctx context.Context) (string, error) {
 *  panic("not implemented")
 *}
 */

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
