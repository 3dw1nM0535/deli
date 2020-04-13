package models

import "github.com/gofrs/uuid"

// Payment : payment data model
type Payment struct {
	BaseModel
	MerchantRequestID string `gorm:"type:varchar(50);not_null;"`
	CheckoutRequestID string `gorm:"type:varchar(50);not_null;"`
	MpesaDescription  string `gorm:"text;not_null;"`
	OrderID           uuid.UUID
	OrderPaidFor      *Order
	RestaurantID      uuid.UUID
	PaidAmount        float64 `gorm:"type:float;not_null;"`
	PhoneNumber       string  `gorm:"type:varchar(50);not_null;"`
	Confirmed         bool    `gorm:"type:boolean;not_null;default:false;"`
}
