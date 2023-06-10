package controllers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type UsersController interface {
	GetUser(c echo.Context) (err error)
}

type UsersControllerParams struct {
	di.Inject

	Users services.UsersService
}

func NewUsersController(params UsersControllerParams) UsersController {
	return &params
}

func (h *UsersControllerParams) GetUser(c echo.Context) (err error) {
	var params dtos.GetUserReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	user, err := h.Users.GetUser(params)
	return responses.New().
		WithData(user).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
