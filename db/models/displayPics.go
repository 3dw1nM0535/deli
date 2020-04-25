package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type DisplayPic struct {
	FileModel
	Media        string `gorm:"text;not_null;"`
	Content      string `gorm:"type:varchar(225);not_null;"`
	Size         int    `gorm:"type:integer;not_null;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RestaurantID uuid.UUID
}
