package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gofrs/uuid"
)

// Dish : menu dish data model
type Dish struct {
	BaseModel
	Title       string         `gorm:"not_null;type:varchar(255);"`
	Description string         `gorm:"not_null;text;"`
	Price       float64        `gorm:"not_null;type:float;"`
	Image       string         `gorm:"not_null;text;"`
	DishAddOn   postgres.Jsonb `gorm:"not_null;"`
	MenuID      uuid.UUID
	AddOns      []*DishAddOn
}
