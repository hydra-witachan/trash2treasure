package repositories

import (
	"go-boilerplate/internal/models"

	"cloud.google.com/go/storage"
	"github.com/goava/di"
	"gorm.io/gorm"
)

type SubCategoriesRepository interface {
	GetSubCategories(categoryID string) (subCategories []models.SubCategory, err error)
}

type SubCategoriesRepositoryParams struct {
	di.Inject

	Bucket *storage.BucketHandle
	Gorm   *gorm.DB
}

func NewSubCategoriesRepository(params SubCategoriesRepositoryParams) SubCategoriesRepository {
	return &params
}

func (r *SubCategoriesRepositoryParams) GetSubCategories(categoryID string) (subCategories []models.SubCategory, err error) {
	err = r.Gorm.Where("category_id = ?", categoryID).Find(&subCategories).Error
	return
}