package util

import (
	"github.com/labstack/echo"
	"net/http"
)

type CustomHTTPSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    interface{} `json:"data"`
}

type CustomHTTPError struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func PublishFailureMessage(message string, c echo.Context) error {
	_error := CustomHTTPError{
		Code:    http.StatusBadRequest,
		Message: message,
		Status:  false,
	}
	return c.JSONPretty(http.StatusBadGateway, _error, "  ")
}

func PublishSuccessData(data interface{}, c echo.Context) error {
	_success := CustomHTTPSuccess{
		Code:    http.StatusBadRequest,
		Message: "Success Get Data",
		Status:  false,
		Data: data,
	}
	return c.JSONPretty(http.StatusBadGateway, _success, "  ")
}
