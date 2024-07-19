package router_v1

import (
	"github.com/labstack/echo/v4"
)

func (r *Router) registerNoSqlRoutes(g *echo.Group) {
	g.GET("/nosql", r.noSqlHandler.GetAll)
	g.GET("/nosql/:id", r.noSqlHandler.GetById)
}
