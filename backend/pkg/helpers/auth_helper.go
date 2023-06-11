package helpers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTString(claims dtos.AuthClaims) (token string, err error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err = rawToken.SignedString(jwtSecret)
	if err != nil {
		err = responses.NewError().
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to sign JWT.")
	}
	return
}
