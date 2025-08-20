package usecases

import (
	"github.com/google/wire"
	"mono-base/internal/usecases/app_version"
	"mono-base/internal/usecases/user"
)

var UserUseCaseProviders = wire.NewSet(
	user.NewGetUserByIdUseCase,
	user.NewLoginUseCase,
	app_version.NewGetListAppVersionUseCase,
)
