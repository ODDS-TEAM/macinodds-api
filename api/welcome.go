package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetWelcome show status 200 and display a welcome message.
func (h *Handler) GetWelcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to mac.odds.team API v.1.0.0 beta!")
}
