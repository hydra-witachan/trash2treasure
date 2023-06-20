package models

import "go-boilerplate/internal/constants"

type User struct {
	BaseModel

	Username string             `gorm:"column:username" json:"username"`
	FullName string             `gorm:"column:full_name" json:"fullName"`
	Email    string             `gorm:"column:email" json:"email"`
	Password string             `gorm:"column:password" json:"-"` // don't allow password to ever be exported.
	Address  string             `gorm:"column:address" json:"address"`
	Role     constants.UserRole `gorm:"column:role" json:"role"`
	Points   int64              `gorm:"column:points" json:"points"`
}

func (User) TableName() string {
	return "users"
}
