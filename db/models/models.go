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
	RestaurantName string `gorm:"unique_index:idx_restaurant;not_null;type:varchar(100)"`
	About          string `gorm:"not_null;type:text"`
	Telephone      string `gorm:"unique_index:idx_telephone;unique;not_null;type:varchar(50)"`
	Verified       bool   `gorm:"unique_index:idx_verfied;default:false;not_null"`
}
