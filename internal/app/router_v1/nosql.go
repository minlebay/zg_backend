package router_v1

import (
	"github.com/labstack/echo/v4"
	"zg_backend/internal/app/handlers"
)

func (r *RouterV1) registerNoSqlRoutes(e *echo.Echo) {
	e.GET("/nosql", handlers.NoSqlRootHandler)
}
