package services

import (
	"context"
	"encoding/base64"
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
	CreateItem(ctx context.Context, claims dtos.AuthClaims, params dtos.CreateItemReq) (newItem models.Item, err error)
	GetItemByID(params dtos.GetItemByIDReq) (item models.Item, err error)
	GetItems(params dtos.GetItemsReq) (items []models.Item, err error)
	DonateItem(claims dtos.AuthClaims, params dtos.DonateItemReq) (err error)
	GetCollectorItems(claims dtos.AuthClaims) (items []models.Item, err error)
}

type ItemsServiceParams struct {
	di.Inject

	UsersService UsersService
	Items        repositories.ItemsRepository
	Users        repositories.UsersRepository
}

func NewItemsService(params ItemsServiceParams) ItemsService {
	return &params
}

func (s *ItemsServiceParams) CreateItem(ctx context.Context, claims dtos.AuthClaims, params dtos.CreateItemReq) (newItem models.Item, err error) {
	user, err := s.UsersService.GetUser(dtos.GetUserReq{UserID: claims.ID})
	if err != nil {
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(params.EncodedImage)
	if err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Cannot decode image from base64.")
		return
	}

	contentType := http.DetectContentType(imageData)
	acceptableTypeMap := map[string]string{
		"image/png":  "png",
		"image/jpeg": "jpg",
	}
	imageFileType, ok := acceptableTypeMap[contentType]
	if !ok {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithMessage("Image data isn't acceptable, make sure it's either PNG or JPEG.")
		return
	}

	if user.Points < int64(params.NeededAmount) * int64(params.PointsPerItem) {
		err = responses.NewError().
			WithError(err).
			WithMessage("not enough point users").
			WithCode(http.StatusBadRequest)
		return
	}
	user.Points -= int64(params.NeededAmount) * int64(params.PointsPerItem)

	newItem = models.Item{
		AuthorID:        claims.ID,
		Category: 	 	 params.Category,
		SubCategory: 	 params.SubCategory,
		AuthorName:      claims.FullName,
		ItemName:        params.ItemName,
		Description:     params.Description,
		Points:          params.PointsPerItem,
		NeededAmount:    params.NeededAmount,
		FullfiledAmount: 0,
		ImageURL:        "", // image url will be filled after it's uploaded.
	}

	err = s.Items.CreateItem(&newItem)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage("Failed to create new item.").
			WithCode(http.StatusInternalServerError)
	}

	imageUrl, err := s.Items.UploadItemImage(ctx, dtos.UploadItemImageParams{
		ItemID:    newItem.ID,
		FileType:  imageFileType,
		ImageData: imageData,
	})
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage("Failed to upload item image.").
			WithCode(http.StatusInternalServerError)
		return
	}

	newItem.ImageURL = imageUrl
	if err = s.Items.UpdateItem(&newItem); err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage("Failed to update item to have image URL.").
			WithCode(http.StatusInternalServerError)
	}
	return
}

func (s *ItemsServiceParams) GetItemByID(params dtos.GetItemByIDReq) (item models.Item, err error) {
	item, err = s.Items.GetItemByID(params.ItemID)
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

func (s *ItemsServiceParams) GetItems(params dtos.GetItemsReq) (items []models.Item, err error) {
	items, err = s.Items.GetItems(params.SubCategory, params.Search)
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get items.")
		err = newErr
	}
	return
}

func (s *ItemsServiceParams) DonateItem(claims dtos.AuthClaims, params dtos.DonateItemReq) (err error) {
	user, err := s.UsersService.GetUser(dtos.GetUserReq{UserID: claims.ID})
	if err != nil {
		return
	}

	item, err := s.GetItemByID(dtos.GetItemByIDReq{ItemID: params.ItemID})
	if err != nil {
		return
	}

	futureTotalAmount := (item.FullfiledAmount + params.Quantity)
	if futureTotalAmount > item.NeededAmount {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithMessage("The donated quantity exceeds the amount needed by the collector.")
		return
	}

	user.Points += int64(item.Points * params.Quantity)
	item.FullfiledAmount = futureTotalAmount

	if err = s.Users.SaveUser(&user); err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to update donator's points.")
		return
	}
	if err = s.Items.UpdateItem(&item); err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to update donated item details (fulfilled amount).")
		return
	}

	return
}

func (s *ItemsServiceParams) GetCollectorItems(claims dtos.AuthClaims) (items []models.Item, err error) {
	items, err = s.Items.GetCollectorItems(claims.ID)
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get items.")
		err = newErr
	}
	return
}