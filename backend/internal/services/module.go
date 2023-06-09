package services

import "github.com/goava/di"

var Module = di.Options(
	di.Provide(NewUsersService),
	di.Provide(NewItemsService),
	di.Provide(NewCategoriesService),
	di.Provide(NewSubCategoriesService),
)
