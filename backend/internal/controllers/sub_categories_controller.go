package controllers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type SubCategoriesController interface {
	GetSubCategories(c echo.Context) (err error)
}

type SubCategoriesControllerParams struct {
	di.Inject

	SubCategories services.SubCategoriesService
}

func NewSubCategoriesController(params SubCategoriesControllerParams) SubCategoriesController {
	return &params
}

func (h *SubCategoriesControllerParams) GetSubCategories(c echo.Context) (err error) {
	var params dtos.GetSubCategoriesReq

	err = c.Bind(&params)
	if err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	categories, err := h.SubCategories.GetSubCategories(params)
	return responses.New().
		WithData(categories).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}