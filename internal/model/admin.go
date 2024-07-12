package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username           string `gorm:"unique;not null" json:"username"`
	Email              string `gorm:"not null" json:"email"`
	Password           string `gorm:"not null" json:"password"`
	ResetPasswordToken string `json:"reset_password_token,omitempty"`
	ResetTokenExpiry   int64  `json:"reset_token_expiry,omitempty"`
}
