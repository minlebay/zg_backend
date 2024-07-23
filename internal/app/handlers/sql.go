package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zg_backend/internal/app/services"
	"zg_backend/internal/model"
)

type SqlHandler struct {
	sqlService services.SqlService
}

func NewSqlHandler(sqlService services.SqlService) *SqlHandler {
	return &SqlHandler{sqlService: sqlService}
}

// GetAll godoc
// @Summary Get all messages from mysql
// @Description Get all messages from mysql
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /sql [get]
// @Param page query int false "Page number"
// @Param size query int false "Page size"
func (h *SqlHandler) GetAll(c echo.Context) error {
	page, size := getPagination(c)
	var messages []*model.Message
	messages, err := h.sqlService.GetAll(page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(messages))

}

// GetById godoc
// @Summary Get a message by ID from mysql
// @Description Get a message by ID from mysql
// @Accept json
// @Produce json
// @Param id path string true "Message ID"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /sql/{id} [get]
func (h *SqlHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	var message *model.Message
	message, err := h.sqlService.GetMessageByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, successResponse(message))
}
