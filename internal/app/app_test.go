package app

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"testing"
	"zg_backend/internal/app/cache"
	"zg_backend/internal/app/nosql_repository"
	"zg_backend/internal/app/server"
	"zg_backend/internal/app/sql_kv_db"
	"zg_backend/internal/app/sql_repository"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(
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
	require.NoError(t, err)
}
