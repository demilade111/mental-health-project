package models

import "gorm.io/gorm"

type UserRole string

const (
	RoleUser      UserRole = "user"
	RoleTherapist UserRole = "therapist"
	RoleAdmin     UserRole = "admin"
)

type User struct {
	gorm.Model
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
	Email              string   `gorm:"unique" json:"email"`
	Password           string   `json:"-" validate:"required,min=8"`
	Role               UserRole `gorm:"type:varchar(20);default:user" json:"role"`
	Verified           bool     `gorm:"default:false" json:"verified"`
	VerificationToken  string   `json:"verification_token"`         
	PasswordResetToken string   `json:"password_reset_token"`       
}
