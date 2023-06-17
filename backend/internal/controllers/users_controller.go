package controllers

import (
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/helpers"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type UsersController interface {
	GetUser(c echo.Context) (err error)
	Register(c echo.Context) (err error)
	Login(c echo.Context) (err error)
	UserTopup(c echo.Context) (err error)
	RedeemPoints(c echo.Context) (err error)
	GetRedeemExchangeRate(c echo.Context) (err error)
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

func (h *UsersControllerParams) Register(c echo.Context) (err error) {
	var params dtos.RegisterUserReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	err = h.Users.Register(params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		Send(c)
}

func (h *UsersControllerParams) Login(c echo.Context) (err error) {
	var params dtos.LoginUserReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	res, err := h.Users.Login(params)
	return responses.New().
		WithData(res).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}

func (h *UsersControllerParams) UserTopup(c echo.Context) (err error) {
	var params dtos.UserTopupReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	claims := c.Get(constants.AuthClaimsKey).(dtos.AuthClaims)

	err = h.Users.UserTopup(params, claims)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}

func (h *UsersControllerParams) RedeemPoints(c echo.Context) (err error) {
	var params dtos.RedeemPointsReq
	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")
		return
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return
	}

	res, err := h.Users.RedeemPoints(claims, params)
	return responses.New().
		WithData(res).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}

func (h *UsersControllerParams) GetRedeemExchangeRate(c echo.Context) (err error) {
	exchangeRate := constants.RedeemExchangeRate
	return responses.New().
		WithData(exchangeRate).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
