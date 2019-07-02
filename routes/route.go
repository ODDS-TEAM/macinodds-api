package route

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/api"
	"gopkg.in/mgo.v2"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	const DBUrl = "localhost:27017"
	db, err := mgo.Dial(DBUrl)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize handler
	api := &api.Handler{
		DB: db,
	}

	// Routes
	e.GET("/", api.GetWelcome)
	// e.POST("/signin", api.SignIn)

	m := e.Group("/mac")
	m.GET("", api.GetMac)
	m.GET("/:id", api.GetMacByID)
	m.POST("", api.CreateMac)
	m.PUT("/:id", api.UpdateMac)
	m.DELETE("/:id", api.RemoveMac)

	// Delete data in directly database
	e.DELETE("/db/:id", api.RemoveMacInDB)
}
