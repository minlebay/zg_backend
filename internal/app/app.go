package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_backend/internal/app/cache"
	"zg_backend/internal/app/nosql_repository"
	"zg_backend/internal/app/server"
	"zg_backend/internal/app/sql_kv_db"
	"zg_backend/internal/app/sql_repository"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			cache.NewModule(),
			sql_kv_db.NewModule(),
			sql_repository.NewModule(),
			nosql_repository.NewModule(),
			server.NewModule(),
		),
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
}
