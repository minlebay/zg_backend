package router_v1

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "zg_backend/docs"
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
	g := e.Group("api/v1")

	r.registerSqlRoutes(g)
	r.registerNoSqlRoutes(g)

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

}
