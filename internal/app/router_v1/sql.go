package router_v1

import (
	"github.com/labstack/echo/v4"
)

func (r *Router) registerSqlRoutes(g *echo.Group) {
	g.GET("/sql", r.sqlHandler.GetAll)
	g.GET("/sql/:id", r.sqlHandler.GetById)
}
