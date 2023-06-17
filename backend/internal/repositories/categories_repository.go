package repositories

import (
	"go-boilerplate/internal/models"

	"cloud.google.com/go/storage"
	"github.com/goava/di"
	"gorm.io/gorm"
)

type CategoriesRepository interface {
	GetCategories() (categories []models.Category, err error)
}

type CategoriesRepositoryParams struct {
	di.Inject

	Bucket *storage.BucketHandle
	Gorm   *gorm.DB
}

func NewCategoriesRepository(params CategoriesRepositoryParams) CategoriesRepository {
	return &params
}

func (r *CategoriesRepositoryParams) GetCategories() (categories []models.Category, err error) {
	err = r.Gorm.Find(&categories).Error
	return
}