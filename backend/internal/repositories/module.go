package repositories

import "github.com/goava/di"

var Module = di.Options(
	di.Provide(NewUsersRepository),
	di.Provide(NewItemsRepository),
	di.Provide(NewCategoriesRepository),
	di.Provide(NewSubCategoriesRepository),
)
