package dtos

import (
	"github.com/labstack/echo/v4"
)

// MediaDto struct define StatusCode, Message, and Data property
// to represent the API response type.
type MediaDto struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}
