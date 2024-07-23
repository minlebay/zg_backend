package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zg_backend/internal/app/services"
	"zg_backend/internal/model"
)

type NoSqlHandler struct {
	noSqlService services.NoSqlService
}

func NewNoSqlHandler(noSqlService services.NoSqlService) *NoSqlHandler {
	return &NoSqlHandler{noSqlService: noSqlService}
}

// GetAll godoc
// @Summary Get all messages from mongodb
// @Description Get all messages from mongodb
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /nosql [get]
// @Param page query int false "Page number"
// @Param size query int false "Page size"
func (h *NoSqlHandler) GetAll(c echo.Context) error {
	page, size := getPagination(c)
	var messages []*model.Message
	messages, err := h.noSqlService.GetAll(page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(messages))
}

// GetById godoc
// @Summary Get a message by ID from mongodb
// @Description Get a message by ID from mongodb
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /nosql/{id} [get]
func (h *NoSqlHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	var message *model.Message
	message, err := h.noSqlService.GetMessageByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(message))
}
