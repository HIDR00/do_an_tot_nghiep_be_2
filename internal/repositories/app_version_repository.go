package repositories

import (
	"context"
	"mono-base/internal/entities"
)

type AppVersionRepository interface {
	GetListAppVersion(ctx context.Context) ([]entities.AppVersionData, error)
}
