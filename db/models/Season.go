package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time
}

type Season struct {
	Model
	Token         int    `gorm:"type:integer;"`
	Crop          string `gorm:"type:varchar(255);"`
	Fertilizer    string `gorm:"type:varchar(255);"`
	Seed          string `gorm:"type:varchar(255);"`
	ExpectedYield string `gorm:"type:varchar(50);"`
	SeedSupplier  string `gorm:"type:varchar(255);"`
	HarvestYield  string `gorm:"type:varchar(50);"`
	HarvestPrice  string `gorm:"type:varchar(255);"`
}
