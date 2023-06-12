package repositories

import (
	"go-boilerplate/internal/models"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type ItemsRepository interface {
	CreateItem(item models.Item) (err error)
}

type ItemsRepositoryParams struct {
	di.Inject

	Gorm *gorm.DB
}

func NewItemsRepository(params ItemsRepositoryParams) ItemsRepository {
	return &params
}

func (r *ItemsRepositoryParams) CreateItem(item models.Item) (err error) {
	err = r.Gorm.Create(&item).Error
	return
}