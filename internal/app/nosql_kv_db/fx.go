package nosql_kv_db

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var _ NosqlKvDb = (*Redis)(nil)

type NosqlKvDb interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
	Get(key string) (out []byte, err error)
	Put(key string, value []byte) (err error)
	Delete(key string) (err error)
	Iterate(filter string) (keys []string, err error)
}

func NewModule() fx.Option {

	return fx.Module(
		"redis",
		fx.Provide(
			NewKeyValueDbConfig,
			fx.Annotate(
				NewRedis,
				fx.As(new(NosqlKvDb)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, c NosqlKvDb) {
				lc.Append(fx.StartStopHook(c.Start, c.Stop))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("redis")
		}),
	)
}
