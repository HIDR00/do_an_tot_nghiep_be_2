package usecases

import (
	"github.com/google/wire"
	"mono-base/internal/usecases/user"
)

var UserUseCaseProviders = wire.NewSet(
	user.NewGetUserByIdUseCase,
	user.NewLoginUseCase,
)
