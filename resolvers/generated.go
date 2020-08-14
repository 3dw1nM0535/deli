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
 *func (r *mutationResolver) UpdateFarmSeason(ctx context.Context, input models.SeasonUpdateInput) (*models1.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmPreparations(ctx context.Context, input models.PreparationInput) (*models1.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmPlantings(ctx context.Context, input models.PlantingInput) (*models1.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmHarvests(ctx context.Context, input *models.HarvestInput) (*models1.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) Hello(ctx context.Context) (string, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) GetFarms(ctx context.Context) ([]*models1.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) GetSeasons(ctx context.Context, input models.SeasonsQueryInput) ([]*models1.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *seasonResolver) ID(ctx context.Context, obj *models1.Season) (string, error) {
 *  panic("not implemented")
 *}
 */

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Season returns graph.SeasonResolver implementation.
func (r *Resolver) Season() graph.SeasonResolver { return &seasonResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
