package handlers

import (
	"github.com/labstack/echo/v4"
)

func SqlRootHandler(c echo.Context) error {
	return c.JSON(200, "Sql Get")
}
