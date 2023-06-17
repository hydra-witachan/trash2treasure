package services

import (
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
)

type CategoriesService interface {
	GetCategories() (categories []models.Category, err error)
}

type CategoriesServiceParams struct {
	di.Inject

	Categories	 repositories.CategoriesRepository
}

func NewCategoriesService(params CategoriesServiceParams) CategoriesService {
	return &params
}

func (s *CategoriesServiceParams) GetCategories() (categories []models.Category, err error) {
	categories, err = s.Categories.GetCategories()
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get categories.")
		err = newErr
	}
	return
}