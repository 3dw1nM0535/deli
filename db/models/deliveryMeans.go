package models

import "github.com/gofrs/uuid"

// DeliveryMeans : means of delivery transprot
type DeliveryMeans struct {
	BaseModel
	Means   string `gorm:"type:varchar(50);not_null;"`
	RiderID uuid.UUID
}
