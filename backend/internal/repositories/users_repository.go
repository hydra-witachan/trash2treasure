package repositories

import (
	"go-boilerplate/internal/models"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUser(id string) (user models.User, err error)
}

type UsersRepositoryParams struct {
	di.Inject

	gorm *gorm.DB
}

func NewUsersRepository(params UsersRepositoryParams) UsersRepository {
	return &params
}

func (r *UsersRepositoryParams) GetUser(id string) (user models.User, err error) {
	err = r.gorm.First(&user, "id = ?", id).Error
	return
}
