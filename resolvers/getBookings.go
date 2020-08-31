package resolvers

import (
	"context"
	"github.com/3dw1nM0535/Byte/db/models"
	models1 "github.com/3dw1nM0535/Byte/models"
)

func (r *queryResolver) GetBookings(ctx context.Context, input models1.BookingsQueryInput) ([]*models.Booking, error) {
	bookings := []*models.Booking{}
	r.ORM.DB.Where("token = ?", input.Token).Find(&bookings)
	return bookings, nil
}
