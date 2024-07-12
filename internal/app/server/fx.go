package server

import (
	"context"
	"go.uber.org/fx"
	"log"
	"time"
)

func NewServer(lc fx.Lifecycle, c *config.Config) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			h.SetOnlineSince(time.Now())
			go e.Start(":8080")
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Println("Stopping server")
			return e.Shutdown(c)
		},
	})

	return e
}
