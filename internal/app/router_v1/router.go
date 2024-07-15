package router_v1

import (
	"context"
	"github.com/labstack/echo/v4"
)

type RouterV1 struct{}

func NewRouter() *RouterV1 {
	return &RouterV1{}
}

func (r *RouterV1) Start(ctx context.Context) {
}

func (r *RouterV1) Stop(ctx context.Context) {
}

func (r *RouterV1) RegisterRoutes(e *echo.Echo) {
	r.registerSqlRoutes(e)
	r.registerNoSqlRoutes(e)
}
