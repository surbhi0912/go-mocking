package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"unique; not null" json:"email" binding:"required, email"`
	Password   string `gorm:"not null" json:"password" binding:"required"`
	SignupType string `json:"signupType"`
}
