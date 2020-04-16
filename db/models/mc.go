package models

import "time"

// MDC : medical certificate data model
type MDC struct {
	FileModel
	Media     string `gorm:"not_null;text;"`
	Content   string `gorm:"type:varchar(255);not_null;"`
	Size      int64  `gorm:"type:integer;not_null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
