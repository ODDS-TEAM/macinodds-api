package route

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/api"
	"gitlab.odds.team/internship/macinodds-api/config"
	"gopkg.in/mgo.v2"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	c := config.Config()

	// Database connection.
	db, err := mgo.Dial(c.DBHost)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize handler
	api := &api.Handler{
		DB: db,
	}

	// Routes
	e.GET("/", api.GetWelcome)

	m := e.Group("/mac")
	m.GET("", api.GetMac)
	m.GET("/:id", api.GetMacByID)

	m.POST("", api.CreateMac)
	m.PUT("/:id", api.UpdateMac)
	m.DELETE("/:id", api.RemoveMac)
}
