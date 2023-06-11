package helpers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTString(claims dtos.AuthClaims) (token string, err error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = rawToken.SigningString()
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to sign JWT.")
	}
	return
}
