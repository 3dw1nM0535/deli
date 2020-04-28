package models

// Restaurant : restaurant data model
type Restaurant struct {
	BaseModel
	RestaurantName string     `gorm:"not_null;type:varchar(100);"`
	About          string     `gorm:"not_null;type:text;"`
	Telephone      string     `gorm:"not_null;type:varchar(50);"`
	Cuisine        string     `gorm:"type:varchar(100);not_null;"`
	Verified       bool       `gorm:"default:false;not_null;"`
	Distance       float64    `gorm:"type:float;default:0;"`
	Address        []*Address `gorm:"many2many:restaurant_addresses;"` // Back-Reference
	License        *License
	RestaurantMenu []*Menu
}
