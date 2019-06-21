package route

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/handler"
	"gopkg.in/mgo.v2"
)

// Init sets
func Init(e *echo.Echo) {
	// Database connection.
	const DBUrl = "139.5.146.213:27017"
	db, err := mgo.Dial(DBUrl)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Initailize routes
	e.GET("/api", h.GetAPI)
	e.GET("/api/devices", h.GetDevices)
	e.GET("/api/devices/:id", h.GetByID)
	e.POST("/api/devices", h.CreateDevice)
	e.PUT("/api/devices/:id", h.UpdateDevice)
	e.DELETE("/api/devices/:id", h.RemoveDevice)
	e.DELETE("/api/db/:id", h.RemoveDB)
}
