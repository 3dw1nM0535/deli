package models

import "github.com/gofrs/uuid"

// Order : order data mode
type Order struct {
	BaseModel
	Notes           []*DishOrder
	RestaurantNotes string `gorm:"text;"`
	OrderStatus     string `gorm:"type:varchar(50);default:'In-Kitchen';not_null;"`
	PaidFor         bool   `gorm:"type:boolean;default:false;not_null;"`
	RestaurantID    uuid.UUID
}

// ID : order id after create
var ID string

// AfterCreate : hook
func (o *Order) AfterCreate() (err error) {
	ID = o.ID.String()
	return
}
