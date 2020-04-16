package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// DisplayPicture : rider dp file data model
type DisplayPicture struct {
	FileModel
	Media     string `gorm:"text;not_null;"`
	Size      int    `gorm:"type:integer;not_null;"`
	Content   string `gorm:"type:varchar(50);not_null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	RiderID   uuid.UUID
}
