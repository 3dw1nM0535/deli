package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) AddBooking(ctx context.Context, input *models1.BookingInput) (*models.Booking, error) {
	newBooking := &models.Booking{
		Volume:    input.Volume,
		Booker:    input.Booker,
		Deposit:   input.Deposit,
		Token:     input.Token,
		Delivered: input.Delivered,
		Cancelled: input.Cancelled,
	}
	r.ORM.DB.Create(&newBooking)
	return newBooking, nil
}
