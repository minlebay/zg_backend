package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zg_backend/internal/app/services"
)

type NoSqlHandler struct {
	noSqlService services.NoSqlService
}

func NewNoSqlHandler(noSqlService services.NoSqlService) *NoSqlHandler {
	return &NoSqlHandler{noSqlService: noSqlService}
}

func (h *NoSqlHandler) GetAll(c echo.Context) error {
	page, size := getPagination(c)
	messages, err := h.noSqlService.GetAll(nil, page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(messages))
}

func (h *NoSqlHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	message, err := h.noSqlService.GetMessageByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(message))
}
