package repository_impl

import (
	"context"
	"mono-base/internal/entities"
	"mono-base/internal/repositories"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type appVersionRepository struct {
	db    *sqlx.DB
	table string
}

func NewAppVersionRepository(db *sqlx.DB) repositories.AppVersionRepository {
	return &appVersionRepository{
		db:    db,
		table: "app_version",
	}
}

func (r *appVersionRepository) GetListAppVersion(ctx context.Context) ([]entities.AppVersionData, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	qb := psql.Select("*").From(r.table)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}
	var appVersions []entities.AppVersionData
	if err := r.db.Select(&appVersions, query, args...); err != nil {
		return nil, err
	}
	return appVersions, nil
}
