package models

import "time"

type Farm struct {
	ID        int       `gorm:"primary_key;not_null;type:integer;"`
	Size      string    `gorm:"not_null;type:varchar(50);"`
	Soil      string    `gorm:"not_null;type:varchar(50);"`
	ImageHash string    `gorm:"not_null;type:text;"`
	Owner     string    `gorm:"type:text;not_null;"`
	Season    string    `gorm:"not_null;type:varchar(100);"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time
}
