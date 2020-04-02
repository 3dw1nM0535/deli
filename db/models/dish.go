package models

import (
	"github.com/lib/pq"

	"github.com/gofrs/uuid"
)

// Dish : menu dish data model
type Dish struct {
	BaseModel
	Title       string         `gorm:"not_null;type:varchar(255);"`
	Description string         `gorm:"not_null;text;"`
	Price       float64        `gorm:"not_null;type:float;"`
	Image       string         `gorm:"not_null;text;"`
	AddOns      pq.StringArray `gorm:"type:varchar(100)[]"`
	MenuID      uuid.UUID
}
