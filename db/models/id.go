package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// IDD : identification document data mode
type IDD struct {
	FileModel
	Media     string `gorm:"not_null;text;"`
	Content   string `gorm:"type:varchar(255);not_null;"`
	Size      int    `gorm:"type:integer;not_null;"`
	RiderID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
