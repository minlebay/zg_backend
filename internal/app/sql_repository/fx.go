package sql_repository

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_backend/internal/model"
)

type SqlRepository interface {
	Start()
	Stop()
	GetAll() ([]*model.Message, error)
	GetById(uuid string) (*model.Message, error)
}

func NewModule() fx.Option {

	return fx.Module(
		"repo",
		fx.Provide(
			NewRepositoryConfig,
			fx.Annotate(
				NewMySQLRepository,
				fx.As(new(SqlRepository)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r SqlRepository) {
				lc.Append(fx.StartStopHook(r.Start, r.Start))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("repo")
		}),
	)
}
