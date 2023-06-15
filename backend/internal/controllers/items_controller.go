package controllers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/helpers"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type ItemsController interface {
	CreateItem(c echo.Context) (err error)
	GetItemByID(c echo.Context) (err error)
	DonateItem(c echo.Context) (err error)
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

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return
	}

	item, err := h.Items.CreateItem(c.Request().Context(), claims, params)
	return responses.New().
		WithData(item).
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		Send(c)
}

func (h *ItemsControllerParams) GetItemByID(c echo.Context) (err error) {
	var params dtos.GetItemByIDReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	item, err := h.Items.GetItemByID(params)
	return responses.New().
		WithData(item).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}

func (h *ItemsControllerParams) DonateItem(c echo.Context) (err error) {
	var params dtos.DonateItemReq

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

	err = h.Items.DonateItem(claims, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
