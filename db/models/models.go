package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type Model struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"index;not_null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Deli struct {
	RestaurantName string         `gorm:"not_null;type:varchar(100)"`
	Telephone      string         `gorm:"not_null;type:varchar(50)"`
	Delicacies     pq.StringArray `gorm:"type:varchar(100)[]"`
	Verified       bool           `gorm:"default:false;not_null"`
	Rating         float64        `gorm:"default:0;not_null"`
	Reviews        pq.StringArray `gorm:"type:varchar(100)[]"`
}
