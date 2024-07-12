package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(

		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),
	)
}
