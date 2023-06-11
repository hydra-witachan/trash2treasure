package models

import "go-boilerplate/internal/constants"

type User struct {
	BaseModel

	Username string             `gorm:"column:username" json:"username"`
	FullName string             `gorm:"full_name" json:"fullName"`
	Email    string             `gorm:"email" json:"email"`
	Password string             `gorm:"password" json:"-"` // don't allow password to ever be exported.
	Role     constants.UserRole `gorm:"role" json:"role"`
}
