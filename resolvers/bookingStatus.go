package resolvers

import (
	"context"
	"errors"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateBookingStatus(ctx context.Context, input models1.BookingStatusInput) (*models.Booking, error) {
	booking := &models.Booking{}
	r.ORM.DB.Where("id = ?", input.ID).Find(&booking)
	if booking.ID.String() == specialUUID && r.ORM.DB.NewRecord(booking) {
		notFound := errors.New("not found")
		return nil, notFound
	}
	booking.Delivered = input.Delivered
	r.ORM.DB.Save(&booking)
	return booking, nil
}
