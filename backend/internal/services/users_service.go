package services

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type UsersService interface {
	GetUser(params dtos.GetUserReq) (user models.User, err error)
}

type UsersServiceParams struct {
	di.Inject

	Users repositories.UsersRepository
}

func NewUsersService(params UsersServiceParams) UsersService {
	return &params
}

func (s *UsersServiceParams) GetUser(params dtos.GetUserReq) (user models.User, err error) {
	// TODO: handle params validation here

	user, err = s.Users.GetUser(params.UserID)
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Failed to get user with ID %s", params.UserID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage(fmt.Sprintf("Cannot find user with ID %s", params.UserID))
		}

		err = newErr
	}

	return
}
