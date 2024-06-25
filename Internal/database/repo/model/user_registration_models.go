package model

// Gender Should be an Enum
type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

// UserRegistration  Gorm Model
type UserRegistrationModel struct {
	Name     string     `gorm:"size:255;not null"`
	Email    string     `gorm:"size:255;primary_key;not null"`
	Password string     `gorm:"size:100;not null"`
	MobileNo string     `gorm:"size:30;primary_key;not null"`
	Gender   GenderType `gorm:"size:8;not null"`
}
