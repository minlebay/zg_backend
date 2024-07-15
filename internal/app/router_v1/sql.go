package router_v1

import (
	"github.com/labstack/echo/v4"
	"zg_backend/internal/app/handlers"
)

func (r *RouterV1) registerSqlRoutes(e *echo.Echo) {
	e.GET("/sql", handlers.SqlRootHandler)
}
