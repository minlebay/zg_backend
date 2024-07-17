package nosql_kv_db

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var _ NosqlKvDb = (*RedisNosqlKvDb)(nil)

type NosqlKvDb interface {
	Start()
	Stop()
	Get(key string) (out []byte, err error)
	Put(key string, value []byte) (err error)
	Delete(key string) (err error)
	Iterate(filter string) (keys []string, err error)
}

func NewModule() fx.Option {

	return fx.Module(
		"nosqlKvDb",
		fx.Provide(
			NewKeyValueDbConfig,
			fx.Annotate(
				NewRedisNosqlKvDb,
				fx.As(new(NosqlKvDb)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, c NosqlKvDb) {
				lc.Append(fx.StartStopHook(c.Start, c.Stop))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("nosqlKvDb")
		}),
	)
}
