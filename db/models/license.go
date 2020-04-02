package models

import (
	"time"
	"github.com/gofrs/uuid"
)

// LicenseModel : override gorm.Model createAt, updatedAt
type LicenseModel struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}

// License : restaurant licence data mode
type License struct {
	LicenseModel
	Media        string `gorm:"not_null;text;"`
	Content      string `gorm:"type:varchar(255);not_null;"`
	Size         int64  `gorm:"type:integer;not_null;"`
	RestaurantID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}