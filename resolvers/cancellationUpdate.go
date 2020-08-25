package resolvers

import (
	"context"
	"errors"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateAfterCancellation(ctx context.Context, input models1.CancellationUpdateInput) (bool, error) {
	season := &models.Season{}
	booking := &models.Booking{}
	r.ORM.DB.Where("token = ? AND season_number = ?", input.Token, input.SeasonNumber).Find(&season)
	if season.ID.String() == specialUUID && r.ORM.DB.NewRecord(season) {
		notFound := errors.New("not found")
		return false, notFound
	}
	r.ORM.DB.Where("id = ?", input.BookingID).Find(&booking)
	if booking.ID.String() == specialUUID && r.ORM.DB.NewRecord(booking) {
		notFound := errors.New("not found")
		return false, notFound
	}
	booking.Volume = input.NewVolume
	booking.Deposit = input.NewDeposit
	r.ORM.DB.Save(&booking)
	season.HarvestYield = input.NewSupply
	r.ORM.DB.Save(&season)
	return true, nil
}
