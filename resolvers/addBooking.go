package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) AddBooking(ctx context.Context, input *models1.BookingInput) (*models.Booking, error) {
	booking := &models.Booking{}
	r.ORM.DB.Where("booker = ?", input.Booker).Find(&booking)
	if booking.ID.String() == specialUUID && r.ORM.DB.NewRecord(booking) {
		newBooking := &models.Booking{
			Volume:    input.Volume,
			Booker:    input.Booker,
			Deposit:   input.Deposit,
			Token:     input.Token,
			Delivered: input.Delivered,
		}
		r.ORM.DB.Create(&newBooking)
		return newBooking, nil
	}
	booking.Volume = input.Volume
	booking.Deposit = input.Deposit
	booking.Delivered = input.Delivered
	r.ORM.DB.Save(&booking)
	return booking, nil
}
