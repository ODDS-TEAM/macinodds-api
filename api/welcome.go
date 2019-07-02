package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetAPI is
func (h *Handler) GetWelcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to mac.odds.team API v.1.0.0 beta!")
}
