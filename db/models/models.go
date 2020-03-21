package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
}

type Deli struct {
	BaseModel
	RestaurantName string `gorm:"not_null;type:varchar(100)"`
	About          string `gorm:"not_null;type:text;"`
	Telephone      string `gorm:"not_null;type:varchar(50)"`
	Verified       bool   `gorm:"default:false;not_null"`
}
