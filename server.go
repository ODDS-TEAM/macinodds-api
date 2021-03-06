package main

import (
	"net/http"

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
	s := config.Spec()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
		middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte("sMJuczqQPYzocl1s6SLj"),
			Skipper: func(c echo.Context) bool {
				// Skip authentication for and login requests
				if c.Path() == "/login" || c.Path() == "/_ah/health" {
					return true
				}
				return false
			},
		}),
	)

	// Respond to API health checks.
	// Indicate the server is healthy.
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "mac.odds.team : ok!")
	})

	// Initialize routes
	route.Init(e)
	// Start server
	e.Logger.Fatal(e.Start(s.APIPort))
}
