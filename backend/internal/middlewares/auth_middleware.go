package middlewares

import (
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"go-boilerplate/pkg/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		authHeaderList, ok := c.Request().Header["Authorization"]
		if !ok || len(authHeaderList) == 0 {
			err = responses.NewError().
				WithError(constants.ErrNotLoggedIn).
				WithCode(http.StatusUnauthorized).
				WithMessage("You don't have the permission.")
			return
		}

		authHeader := authHeaderList[0]
		bearerPrefix := "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			err = responses.NewError().
				WithError(constants.ErrNotLoggedIn).
				WithCode(http.StatusUnauthorized).
				WithMessage("Invalid authorization header.")
			return
		}

		token := strings.Replace(authHeader, bearerPrefix, "", 1)
		claims, err := helpers.ParseAndValidateJWT(token)
		if err != nil {
			return
		}

		c.Set(constants.AuthClaimsKey, claims)
		err = next(c)
		return
	}
}
