package models

type RegisteringUser struct {
	Email          string `json:"Email" validate:"required,email"`
	Password       string `json:"Password" validate:"required,min=8,max=15"`
	PasswordSubmit string `json:"Password-Submit" validate:"required,min=8,max=15"`
	Checkbox       bool   `json:"Checkbox" validate:"required"`
}
