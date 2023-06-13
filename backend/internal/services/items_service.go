package services

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
)

type ItemsService interface {
	CreateItem(params dtos.CreateItemReq,  claims dtos.AuthClaims) (err error)
	GetItemByID(params dtos.GetItemByIDReq) (item models.Item, err error)
}

type ItemsServiceParams struct {
	di.Inject

	Items repositories.ItemsRepository
	Users repositories.UsersRepository
}

func NewItemsService(params ItemsServiceParams) ItemsService {
	return &params
}

func (s *ItemsServiceParams) CreateItem(params dtos.CreateItemReq, claims dtos.AuthClaims) (err error) {
	user, err := s.Users.GetUser(dtos.GetUserParams{ ID: claims.ID })
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
	}

	if user.Points < int64(params.NeededAmount) * int64(params.PointsPerItem) {
		err = responses.NewError().
			WithError(err).
			WithMessage("not enough point users").
			WithCode(http.StatusBadRequest)
		return
	}

	newItem := models.Item{
		AuthorID: claims.ID,
		AuthorName: claims.FullName,
		ItemName: params.ItemName,
		Description: params.Description,
		Points: params.PointsPerItem,
		ImageURL: params.ImageURL,
		NeededAmount: params.NeededAmount,
		FullfiledAmount: 0,
	}

	err = s.Items.CreateItem(newItem);
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
	}

	return err
}

func (s *ItemsServiceParams) GetItemByID(params dtos.GetItemByIDReq) (item models.Item, err error) {
	item, err = s.Items.GetItemByID(params.ItemID)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
	}

	return
}