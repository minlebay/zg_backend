package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zg_backend/internal/app/services"
)

type SqlHandler struct {
	sqlService services.SqlService
}

func NewSqlHandler(sqlService services.SqlService) *SqlHandler {
	return &SqlHandler{sqlService: sqlService}
}

func (h *SqlHandler) GetAll(c echo.Context) error {
	page, size := getPagination(c)
	messages, err := h.sqlService.GetAll(page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(messages))

}

func (h *SqlHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	message, err := h.sqlService.GetMessageByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(message))
}
