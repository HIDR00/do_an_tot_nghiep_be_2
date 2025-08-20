package database

import (
	"mono-base/internal/infrastructure/database/postgres"
	"mono-base/internal/infrastructure/database/repository_impl"

	"github.com/google/wire"
)

var DBProvider = wire.NewSet(
	postgres.GetDBContext,
	postgres.NewPSQLMigration,
	postgres.NewUserRepository,
	repository_impl.NewAppVersionRepository,
)
