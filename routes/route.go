package route

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/handler"
)

func Init(e *echo.Echo) {
	// Routes
	e.GET("/api", handler.GetAPI)
	e.GET("/api/devices", handler.GetDevices)
}
