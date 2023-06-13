package services

import (
	"errors"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type ItemsService interface {
	CreateItem(params dtos.CreateItemReq, claims dtos.AuthClaims) (err error)
	GetItem(params dtos.GetItemReq) (item models.Item, err error)
}

type ItemsServiceParams struct {
	di.Inject

	Items repositories.ItemsRepository
}

func NewItemsService(params ItemsServiceParams) ItemsService {
	return &params
}

func (s *ItemsServiceParams) CreateItem(params dtos.CreateItemReq, claims dtos.AuthClaims) (err error) {
	newItem := models.Item{
		AuthorID:        claims.ID,
		AuthorName:      claims.FullName,
		ItemName:        params.ItemName,
		Description:     params.Description,
		Points:          0,
		ImageURL:        "https://google.com",
		NeededAmount:    params.NeededAmount,
		FullfiledAmount: 0,
	}

	err = s.Items.CreateItem(newItem)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
	}

	return err
}

func (s *ItemsServiceParams) GetItem(params dtos.GetItemReq) (item models.Item, err error) {
	item, err = s.Items.GetItem(params.ItemID)
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get item.")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage("Cannot find item.")
		}

		err = newErr
	}

	return
}
