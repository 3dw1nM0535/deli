package models

import (
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

// DishOrder : ordered dish struct
type DishOrder struct {
	BaseModel
	DishID      uuid.UUID
	Title       string         `gorm:"type:varchar(100);not_null;"`
	Description string         `gorm:"text;not_null;"`
	AddOns      pq.StringArray `gorm:"type:varchar(100)[]"`
	Price       float64        `gorm:"type:float;not_null;default:0;"`
	Count       int            `gorm:"type:integer;not_null;default:0;"`
	Subtotal    float64        `gorm:"type:float;not_null;default:0;"`
	OrderID     uuid.UUID
}
