package models

import (
	"github.com/gofrs/uuid"
)

type Address struct {
	BaseModel
	PostalCode    string        `gorm:"not_null;type:varchar(50);"`
	PostalTown    string        `gorm:"not_null;type:varchar(100);"`
	City          string        `gorm:"type:varchar(100);not_null;"`
	BuildingName  string        `gorm:"type:varchar(100);not_null;"`
	StreetName    string        `gorm:"type:varchar(100);not_null;"`
	Lon           float64       `gorm:"type:float;not_null;"`
	Lat           float64       `gorm:"type:float;not_null;"`
	AddressString string        `gorm:"text;"`
	Restaurants   []*Restaurant `gorm:"many2many:restaurant_addresses;"` // Back-Reference
}

// RestaurantAddresses : will hold many2many relationship
type RestaurantAddresses struct {
	RestaurantID uuid.UUID
	AddressID    uuid.UUID
}
