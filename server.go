package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gitlab.odds.team/internship/macinodds-api/config"
	route "gitlab.odds.team/internship/macinodds-api/routes"
)

func main() {
	// Use labstack/echo for rich routing.
	// See https://echo.labstack.com/
	e := echo.New()
	c := config.Config()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
	)

	// Initialize routes
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(c.APIPort))
}
