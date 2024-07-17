package sql_kv_db

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var _ SqlKvDb = (*Redis)(nil)

type SqlKvDb interface {
	Start()
	Stop()
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
				fx.As(new(SqlKvDb)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, c SqlKvDb) {
				lc.Append(fx.StartStopHook(c.Start, c.Stop))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("redis")
		}),
	)
}
