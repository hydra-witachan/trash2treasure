package models

import "go-boilerplate/internal/constants"

type User struct {
	BaseModel

	Username string             `json:"username"`
	FullName string             `json:"fullName"`
	Email    string             `json:"email"`
	Password string             `json:"-"` // don't allow password to ever be exported.
	Address  string             `json:"address"`
	Role     constants.UserRole `json:"role"`
}

func (User) TableName() string {
	return "users"
}
