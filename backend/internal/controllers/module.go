package controllers

import "github.com/goava/di"

var Module = di.Options(
	di.Provide(NewUsersController),
	di.Provide(NewItemsController),
	di.Provide(NewCategoriesController),
	di.Provide(NewSubCategoriesController),
)
