package controllers

import (
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type CategoriesController interface {
	GetCategories(c echo.Context) (err error)
}

type CategoriesControllerParams struct {
	di.Inject

	Categories services.CategoriesService
}

func NewCategoriesController(params CategoriesControllerParams) CategoriesController {
	return &params
}

func (h *CategoriesControllerParams) GetCategories(c echo.Context) (err error) {
	categories, err := h.Categories.GetCategories()
	return responses.New().
		WithData(categories).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}