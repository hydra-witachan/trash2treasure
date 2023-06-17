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
	UserTopup(params dtos.UserTopupReq, claims dtos.AuthClaims) (err error)
	RedeemPoints(claims dtos.AuthClaims, params dtos.RedeemPointsReq) (res dtos.RedeemPointsRes, err error)
}

type UsersServiceParams struct {
	di.Inject

	Users repositories.UsersRepository
}

func NewUsersService(params UsersServiceParams) UsersService {
	return &params
}

func (s *UsersServiceParams) GetUser(params dtos.GetUserReq) (user models.User, err error) {
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
		Role:     string(user.Role),
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

func (s *UsersServiceParams) UserTopup(params dtos.UserTopupReq, claims dtos.AuthClaims) (err error) {
	user, err := s.Users.GetUser(dtos.GetUserParams{ID: claims.ID})
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
		return
	}

	if user.Role != constants.CollectorRole {
		err = responses.NewError().
			WithError(err).
			WithMessage("user role is not a collector").
			WithCode(http.StatusBadRequest)
		return
	}

	user.Points += int64(params.Points)

	err = s.Users.SaveUser(&user)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithMessage(err.Error()).
			WithCode(http.StatusInternalServerError)
	}

	return
}

func (s *UsersServiceParams) RedeemPoints(claims dtos.AuthClaims, params dtos.RedeemPointsReq) (res dtos.RedeemPointsRes, err error) {
	user, err := s.GetUser(dtos.GetUserReq{UserID: claims.ID})
	if err != nil {
		return
	}

	if user.Role != constants.DonatorRole {
		err = responses.NewError().
			WithError(err).
			WithMessage("User's role is not a donator.").
			WithCode(http.StatusBadRequest)
		return
	}

	moneyReceived, ok := constants.RedeemExchangeRate[params.PointsToExchange]
	if !ok {
		err = responses.NewError().
			WithMessage("Invalid points to be exchanged as it doesn't follow the exchange rate maps.").
			WithCode(http.StatusBadRequest)
		return
	}
	if user.Points < params.PointsToExchange {
		err = responses.NewError().
			WithMessage("User doesn't have enough points to be exchanged with real money.").
			WithCode(http.StatusBadRequest)
		return
	}

	res.MoneyReceived = moneyReceived
	user.Points -= params.PointsToExchange

	err = s.Users.SaveUser(&user)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to take user's points for exchange.")
	}
	return
}
