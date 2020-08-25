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
 *func (r *bookingResolver) ID(ctx context.Context, obj *models.Booking) (string, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) AddFarm(ctx context.Context, input models1.FarmInput) (*models.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmSeason(ctx context.Context, input models1.SeasonUpdateInput) (*models.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmPreparations(ctx context.Context, input models1.PreparationInput) (*models.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmPlantings(ctx context.Context, input models1.PlantingInput) (*models.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmHarvests(ctx context.Context, input *models1.HarvestInput) (*models.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateFarmHarvestSupply(ctx context.Context, input models1.HarvestUpdateInput) (*models.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) AddBooking(ctx context.Context, input *models1.BookingInput) (*models.Booking, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *mutationResolver) UpdateAfterCancellation(ctx context.Context, input models1.CancellationUpdateInput) (bool, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) Hello(ctx context.Context) (string, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) GetFarms(ctx context.Context) ([]*models.Farm, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) GetSeasons(ctx context.Context, input models1.SeasonsQueryInput) ([]*models.Season, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *queryResolver) GetBookings(ctx context.Context, input models1.BookingsQueryInput) ([]*models.Booking, error) {
 *  panic("not implemented")
 *}
 *
 *func (r *seasonResolver) ID(ctx context.Context, obj *models.Season) (string, error) {
 *  panic("not implemented")
 *}
 */

// Booking returns graph.BookingResolver implementation.
func (r *Resolver) Booking() graph.BookingResolver { return &bookingResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Season returns graph.SeasonResolver implementation.
func (r *Resolver) Season() graph.SeasonResolver { return &seasonResolver{r} }

type bookingResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
