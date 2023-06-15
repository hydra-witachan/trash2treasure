package main

import (
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/middlewares"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RoutesParams struct {
	di.Inject

	Echo  *echo.Echo
	Users controllers.UsersController
	Items controllers.ItemsController
}

func SetupRoutes(p RoutesParams) {
	usersGroup := p.Echo.Group("users")
	usersGroup.GET("/:id", p.Users.GetUser, middlewares.AuthMiddleware)
	usersGroup.POST("/register", p.Users.Register)
	usersGroup.POST("/login", p.Users.Login)
	usersGroup.PATCH("/topup", p.Users.UserTopup, middlewares.AuthMiddleware)
	usersGroup.POST("/redeem", p.Users.RedeemPoints, middlewares.AuthMiddleware)
	usersGroup.GET("/redeem-exchange-rate", p.Users.GetRedeemExchangeRate)

	itemsGroup := p.Echo.Group("items")
	itemsGroup.POST("", p.Items.CreateItem, middlewares.AuthMiddleware)
	itemsGroup.GET("/:id", p.Items.GetItemByID)
	itemsGroup.GET("", p.Items.GetItems)
	itemsGroup.POST("/donate/:id", p.Items.DonateItem, middlewares.AuthMiddleware)
}
