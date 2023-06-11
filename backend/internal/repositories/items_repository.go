package repositories

import (
	"github.com/goava/di"
	"gorm.io/gorm"
)

type ItemsRepository interface {
}

type ItemsRepositoryParams struct {
	di.Inject

	gorm *gorm.DB
}

func NewItemsRepository(params ItemsRepositoryParams) ItemsRepository {
	return &params
}
