package router_v1

import (
	"go.uber.org/fx"
)

func NewModule() fx.Option {

	return fx.Module(
		"router",
		fx.Provide(
			fx.Annotate(
				NewRouter,
				fx.As(new(RouterV1)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r *RouterV1) {
				lc.Append(fx.StartStopHook(r.Start, r.Stop))
			},
		),
	)
}
