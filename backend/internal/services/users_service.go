package services

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/pkg/helpers"
	"go-boilerplate/pkg/responses"
	"net/http"
	"time"

	"github.com/goava/di"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersService interface {
	GetUser(params dtos.GetUserReq) (user models.User, err error)
	Register(params dtos.RegisterUserReq) (err error)
	Login(params dtos.LoginUserReq) (res dtos.LoginUserRes, err error)
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

	user, err = s.Users.GetUser(dtos.GetUserParams{
		ID: params.UserID,
	})
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

	newUser := models.User{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Address:  params.Address,
		Role:     params.Role,
		Password: params.Password,
	}
	if err = s.Users.Register(newUser); err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage("Failed to register new user.").
			WithCode(http.StatusInternalServerError)
	}
	return
}

func (s *UsersServiceParams) Login(params dtos.LoginUserReq) (res dtos.LoginUserRes, err error) {
	user, err := s.Users.GetUser(dtos.GetUserParams{
		Email: params.Email,
	})
	if err != nil {
		newErr := responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get user.")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage("Cannot find user.")
		}

		err = newErr
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(params.Password)); err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusUnauthorized).
			WithMessage("Incorrect password.")

		return
	}

	tokenExpireDuration := (time.Hour * 24) // one day
	currentTime := time.Now()

	token, err := helpers.GenerateJWTString(dtos.AuthClaims{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	})
	if err != nil {
		return
	}

	res.AccessToken = token
	return
}
