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
