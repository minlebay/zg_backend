package router_v1

import (
	"github.com/labstack/echo/v4"
)

func (r *Router) registerSqlRoutes(e *echo.Echo) {
	e.GET("/sql", r.sqlHandler.GetAll)
	e.GET("/sql/:id", r.sqlHandler.GetById)
}
