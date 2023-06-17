package services

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
)

type SubCategoriesService interface {
	GetSubCategories(params dtos.GetSubCategoriesReq) (subCategories []models.SubCategory, err error)
}

type SubCategoriesServiceParams struct {
	di.Inject

	SubCategories repositories.SubCategoriesRepository
}

func NewSubCategoriesService(params SubCategoriesServiceParams) SubCategoriesService {
	return &params
}

func (s *SubCategoriesServiceParams) GetSubCategories(params dtos.GetSubCategoriesReq) (subCategories []models.SubCategory, err error) {
	subCategories, err = s.SubCategories.GetSubCategories(params.CategoryID)
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get sub categories.")
		err = newErr
	}
	return
}