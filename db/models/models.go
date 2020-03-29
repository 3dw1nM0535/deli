package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// BaseModel : override gorm.Model
type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
}

// Restaurant : restaurant data model
type Restaurant struct {
	BaseModel
	RestaurantName string     `gorm:"unique_index:idx_restaurant;not_null;type:varchar(100);"`
	About          string     `gorm:"not_null;type:text;"`
	Telephone      string     `gorm:"unique_index:idx_telephone;unique;not_null;type:varchar(50);"`
	Verified       bool       `gorm:"unique_index:idx_verfied;default:false;not_null;"`
	Address        []*Address `gorm:"many2many:restaurant_addresses;"` // Back-Reference
}

// Address : address data model
type Address struct {
	BaseModel
	PostalCode   string        `gorm:"index:idx_postal_no;not_null;type:varchar(50);"`
	PostalTown   string        `gorm:"unique_index:idx_town;not_null;type:varchar(100);"`
	BuildingName string        `gorm:"type:varchar(100);not_null;"`
	StreetName   string        `gorm:"index:idx_street;type:varchar(100);not_null;"`
	Lon          float64       `gorm:"type:float;not_null;"`
	Lat          float64       `gorm:"type:float;not_null;"`
	Restaurants  []*Restaurant `gorm:"many2many:restaurant_addresses;"` // Back-Reference
}
