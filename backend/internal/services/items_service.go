package services

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/repositories"

	"github.com/goava/di"
)

type ItemsService interface {
	CreateItem(body dtos.CreateItemReq) (err error)
}

type ItemsServiceParams struct {
	di.Inject

	Items repositories.ItemsRepository
}

func NewItemsService(params ItemsServiceParams) ItemsService {
	return &params
}

func (s *ItemsServiceParams) CreateItem(body dtos.CreateItemReq) (err error) {
	return err
}