package hander

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorJSON(c echo.Context, status int, code, msg string) error {
	return c.JSON(status, ErrorResponse{Code: code, Message: msg})
}
