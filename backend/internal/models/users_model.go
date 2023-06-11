package models

import "go-boilerplate/internal/constants"

type User struct {
	BaseModel

	Username string             `json:"username"`
	FullName string             `json:"fullName"`
	Email    string             `json:"email"`
	Password string             `json:"-"` // don't allow password to ever be exported.
	Role     constants.UserRole `json:"role"`
}
