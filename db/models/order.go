package models

import "github.com/gofrs/uuid"

// Order : order data mode
type Order struct {
	BaseModel
	Notes           []*DishOrder
	RestaurantNotes string `gorm:"text;"`
	RestaurantID    uuid.UUID
}

// ID : order id after create
var ID string

// AfterCreate : hook
func (o *Order) AfterCreate() (err error) {
	ID = o.ID.String()
	return
}
