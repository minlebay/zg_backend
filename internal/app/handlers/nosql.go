package handlers

import (
	"github.com/labstack/echo/v4"
)

func NoSqlRootHandler(c echo.Context) error {
	return c.JSON(200, "NoSql Get")
}
