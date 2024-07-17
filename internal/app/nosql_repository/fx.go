package nosql_repository

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_backend/internal/model"
)

type NoSqlRepository interface {
	Start()
	Stop()
	GetMessages(filter interface{}) ([]*model.Message, error)
	GetById(id string) (*model.Message, error)
}

func NewModule() fx.Option {

	return fx.Module(
		"repo",
		fx.Provide(
			NewNoSqlRepositoryConfig,
			fx.Annotate(
				NewMongoRepository,
				fx.As(new(NoSqlRepository)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r NoSqlRepository) {
				lc.Append(fx.StartStopHook(r.Start, r.Start))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("repo")
		}),
	)
}
