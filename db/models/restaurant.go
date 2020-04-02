package models

// Restaurant : restaurant data model
type Restaurant struct {
	BaseModel
	RestaurantName string     `gorm:"not_null;type:varchar(100);"`
	About          string     `gorm:"not_null;type:text;"`
	Telephone      string     `gorm:"not_null;type:varchar(50);"`
	Verified       bool       `gorm:"default:false;not_null;"`
	Address        []*Address `gorm:"many2many:restaurant_addresses;"` // Back-Reference
	License        *License
	RestaurantMenu []*Menu
}
