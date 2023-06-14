package helpers

import (
	"errors"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func ParseAndValidateJWT(token string) (claims dtos.AuthClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("Failed to validate token.")
	}

	return
}

func GetAuthClaims(ctx echo.Context) (claims dtos.AuthClaims, err error) {
	claims, ok := ctx.Get(constants.AuthClaimsKey).(dtos.AuthClaims)
	if !ok {
		err = responses.NewError().
			WithCode(http.StatusInternalServerError).
			WithError(errors.New("failed to cast context value to user's claims")).
			WithMessage("Failed to get user's JWT claims.")
	}
	return
}
