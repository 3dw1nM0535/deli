package models

type Booking struct {
	Model
	Volume    int    `gorm:"type:integer;"`
	Booker    string `gorm:"type:varchar(255);"`
	Deposit   string `gorm:"type:varchar(255);"`
	Token     int    `gorm:"type:integer;"`
	Delivered bool   `gorm:"type:boolean;"`
	Cancelled bool   `gorm:"type:boolean;"`
}
