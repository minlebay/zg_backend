package router_v1

import (
	"github.com/labstack/echo/v4"
)

func (r *Router) registerNoSqlRoutes(e *echo.Echo) {
	e.GET("/nosql", r.noSqlHandler.GetAll)
	e.GET("/nosql/:id", r.noSqlHandler.GetById)
}
