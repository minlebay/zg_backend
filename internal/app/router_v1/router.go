package router_v1

import (
	"github.com/labstack/echo/v4"
	"zg_backend/internal/app/handlers"
)

type Router struct {
	sqlHandler   *handlers.SqlHandler
	noSqlHandler *handlers.NoSqlHandler
}

func NewRouter(sqlHandler *handlers.SqlHandler, noSqlHandler *handlers.NoSqlHandler) *Router {
	return &Router{
		sqlHandler:   sqlHandler,
		noSqlHandler: noSqlHandler,
	}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	r.registerSqlRoutes(e)
	r.registerNoSqlRoutes(e)
}
