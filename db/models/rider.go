package models

// Rider : rider data model
type Rider struct {
	BaseModel
	FirstName              string `gorm:"type:varchar(50);not_null;"`
	LastName               string `gorm:"type:varchar(50);not_null;"`
	PhoneNumber            string `gorm:"type:varchar(50);not_null;"`
	EmailAddress           string `gorm:"type:varchar(100);not_null;"`
	Verified               bool   `gorm:"type:boolean;not_null;default:false;"`
	IdentificationDocument *IDD
	MedicalCertificate     *MDC
	GoodConductCertificate *GCC
	DisplayPic             *DisplayPicture
}
