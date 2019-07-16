package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	route "gitlab.odds.team/internship/macinodds-api/routes"
)

func main() {
	// Use labstack/echo for rich routing.
	// See https://echo.labstack.com/
	e := echo.New()
	// s := config.Spec()

	// Respond to API health checks.
	// Indicate the server is healthy.
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "mac.odds.team : ok!")
	})

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
	e.Logger.Fatal(e.Start(":1323"))
}
