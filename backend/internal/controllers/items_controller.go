package controllers

import (
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type ItemsController interface {
	CreateItem(c echo.Context) (err error)
}

type ItemsControllerParams struct {
	di.Inject

	Items services.ItemsService
}

func NewItemsController(params ItemsControllerParams) ItemsController {
	return &params
}

func (h *ItemsControllerParams) CreateItem(c echo.Context) (err error) {
	var params dtos.CreateItemReq

	err = c.Bind(&params)
	if err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	claims := c.Get(constants.AuthClaimsKey).(dtos.AuthClaims)

	err = h.Items.CreateItem(params, claims)
	return responses.New().
	WithError(err).
	WithSuccessCode(http.StatusCreated).
	Send(c)
}