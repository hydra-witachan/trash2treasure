package repositories

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUser(params dtos.GetUserParams) (user models.User, err error)
	Register(model models.User) (err error)
	IsUserExists(params dtos.IsUserExistsParams) (isExists bool, err error)
	SaveUser(model *models.User) (err error)
}

type UsersRepositoryParams struct {
	di.Inject

	Gorm *gorm.DB
}

func NewUsersRepository(params UsersRepositoryParams) UsersRepository {
	return &params
}

func (r *UsersRepositoryParams) GetUser(params dtos.GetUserParams) (user models.User, err error) {
	query := r.Gorm
	if params.ID != "" {
		query.Where("id = ?", params.ID)
	}
	if params.Email != "" {
		query.Where("email = ?", params.Email)
	}

	err = query.Find(&user).Error
	return
}

func (r *UsersRepositoryParams) Register(model models.User) (err error) {
	err = r.Gorm.Create(&model).Error
	return
}

func (r *UsersRepositoryParams) IsUserExists(params dtos.IsUserExistsParams) (isExists bool, err error) {
	var user models.User

	query := r.Gorm
	if params.Email != "" {
		query.Where("email = ?", params.Email)
	}
	if params.Username != "" {
		query.Where("username = ?", params.Username)
	}

	if err = query.Limit(1).Find(&user).Error; err != nil {
		return
	}

	isExists = (user.ID != "")
	return
}

func (r *UsersRepositoryParams) SaveUser(user *models.User) (err error) {
	err = r.Gorm.Save(user).Error
	return
}
