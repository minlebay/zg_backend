package nosql_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_backend/internal/model"
)

type NoSqlRepository interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
	GetAll(ctx context.Context, db mongo.Database) ([]*model.Message, error)
	Create(ctx context.Context, db mongo.Database, entity *model.Message) (*model.Message, error)
	GetById(ctx context.Context, db mongo.Database, id string) (*model.Message, error)
	Update(ctx context.Context, db mongo.Database, id string, entity *model.Message) (*model.Message, error)
	Delete(ctx context.Context, db mongo.Database, id string) error
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
