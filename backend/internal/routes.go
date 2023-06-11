package main

import (
	"go-boilerplate/internal/controllers"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RoutesParams struct {
	di.Inject

	Echo  *echo.Echo
	Users controllers.UsersController
}

func SetupRoutes(p RoutesParams) {
	usersGroup := p.Echo.Group("users")
	// itemsGroup := p.Echo.Group("items")

	usersGroup.GET("/:id", p.Users.GetUser)
	usersGroup.POST("/register", p.Users.Register)
	usersGroup.POST("/login", p.Users.Login)
}
