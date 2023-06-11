package services

import "github.com/goava/di"

var Module = di.Options(
	di.Provide(NewUsersService),
)
