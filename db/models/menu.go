package models

import "github.com/gofrs/uuid"

// Menu : restaurant menu data model
type Menu struct {
	BaseModel
	Headline     string `gorm:"not_null;type:varchar(255);"`
	Dishes       []*Dish
	RestaurantID uuid.UUID
}
