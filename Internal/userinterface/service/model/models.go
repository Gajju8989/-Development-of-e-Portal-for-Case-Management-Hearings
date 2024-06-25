package model

type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

type UserRegistrationRequest struct {
	Name     string     `json:"name" binding:"required"`
	Email    string     `json:"email" binding:"required"`
	Password string     `json:"password" binding:"required"`
	MobileNo string     `json:"mobile_no" binding:"required"`
	Gender   GenderType `json:"gender" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
