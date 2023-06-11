package services

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/goava/di"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersService interface {
	GetUser(params dtos.GetUserReq) (user models.User, err error)
	Register(params dtos.RegisterUserReq) (err error)
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

func (s *UsersServiceParams) Register(params dtos.RegisterUserReq) (err error) {
	passBytes := []byte(params.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, constants.DefaultHashCost)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithMessage("Failed to hash password.").
			WithCode(http.StatusInternalServerError)
	}
	params.Password = string(hashedPassword)

	isExists, err := s.Users.IsUserExists(dtos.IsUserExistsParams{
		Email:    params.Email,
		Username: params.Username,
	})
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithMessage("Failed to check user existence.").
			WithCode(http.StatusInternalServerError)
	}
	if isExists {
		return responses.NewError().
			WithMessage("User with this email or username has already been registered.").
			WithCode(http.StatusBadRequest)
	}

	if err = s.Users.Register(params.User); err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage("Failed to register new user.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
