package dtos

import (
	"go-boilerplate/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserReq struct {
	models.User

	ID       string `json:"-"` // ignore inserting ID
	Password string `json:"password"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	AccessToken string `json:"accessToken"`
}

type GetUserReq struct {
	UserID string `param:"id"`
}

type GetUserParams struct {
	ID    string
	Email string
}

type IsUserExistsParams struct {
	Email    string
	Username string
}

type AuthClaims struct {
	jwt.RegisteredClaims

	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
}

type UserTopupReq struct {
	Points int    `json:"points"`
	Fee    int64  `json:"fee"`
	Method string `json:"method"`
}

type RedeemPointsReq struct {
	PointsToExchange int64 `json:"pointsToExchange"`
}

type RedeemPointsRes struct {
	MoneyReceived int64 `json:"moneyReceived"`
}
