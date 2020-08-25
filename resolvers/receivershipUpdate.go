package resolvers

import (
	"context"
	"errors"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateAfterReceivership(ctx context.Context, input models1.ReceivershipUpdateInput) (bool, error) {
	booking := &models.Booking{}
	r.ORM.DB.Where("id = ?", input.BookingID).Find(&booking)
	if booking.ID.String() == specialUUID && r.ORM.DB.NewRecord(booking) {
		notFound := errors.New("not found")
		return false, notFound
	}
	booking.Volume = input.NewBookerVolume
	booking.Deposit = input.NewBookerDeposit
	booking.Delivered = input.Delivered
	r.ORM.DB.Save(&booking)
	return true, nil
}
