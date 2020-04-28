package models

import (
	"github.com/gofrs/uuid"
)

type DishAddOn struct {
	BaseModel
	Name   string  `gorm:"type:varchar(100);not_null;"`
	Price  float64 `gorm:"type:float;not_null;"`
	DishID uuid.UUID
}
