package repositories

import (
	"context"
	"fmt"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"log"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/goava/di"
	"gorm.io/gorm"
)

type ItemsRepository interface {
	CreateItem(item *models.Item) (err error)
	UpdateItem(item *models.Item) (err error)
	GetItemByID(id string) (item models.Item, err error)
	GetItems(subCategoryId string, search string) (items []models.Item, err error)
	UploadItemImage(ctx context.Context, params dtos.UploadItemImageParams) (imageUrl string, err error)
}

type ItemsRepositoryParams struct {
	di.Inject

	Bucket *storage.BucketHandle
	Gorm   *gorm.DB
}

func NewItemsRepository(params ItemsRepositoryParams) ItemsRepository {
	return &params
}

func (r *ItemsRepositoryParams) CreateItem(item *models.Item) (err error) {
	err = r.Gorm.Create(item).Error
	return
}

func (r *ItemsRepositoryParams) UpdateItem(item *models.Item) (err error) {
	err = r.Gorm.Save(item).Error
	return
}

func (r *ItemsRepositoryParams) UploadItemImage(ctx context.Context, params dtos.UploadItemImageParams) (imageUrl string, err error) {
	outputFileName := fmt.Sprintf("items/%s.%s", params.ItemID, params.FileType)

	log.Println("Expected output:", outputFileName)

	object := r.Bucket.Object(outputFileName)
	writer := object.NewWriter(ctx)
	defer writer.Close()

	_, err = writer.Write(params.ImageData)
	if err != nil {
		return
	}

	imageUrl = fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media",
		os.Getenv("FIREBASE_BUCKET"),
		url.PathEscape(outputFileName),
	)
	return
}

func (r *ItemsRepositoryParams) GetItemByID(id string) (item models.Item, err error) {
	err = r.Gorm.Where("id = ?", id).First(&item).Error
	return
}

func (r *ItemsRepositoryParams) GetItems(subCategory string, search string) (items []models.Item, err error) {
	query := `SELECT *
		FROM items i 
		WHERE i.sub_category LIKE ? 
	`

	if search != "" {
		query +=fmt.Sprintf("AND i.item_name LIKE '%%%s%%'", search)
	}

	err = r.Gorm.Raw(query, subCategory).Scan(&items).Error
	return
}