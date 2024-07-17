package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func getPagination(c echo.Context) (int, int) {
	pageQ := c.QueryParam("page")
	sizeQ := c.QueryParam("size")

	page, err := strconv.Atoi(pageQ)
	if err != nil {
		page = 0
	}

	size, err := strconv.Atoi(sizeQ)
	if err != nil {
		size = 10
	}

	return page, size
}

func errorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}

func successResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}
