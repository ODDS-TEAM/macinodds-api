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
	s := config.Spec()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
		// middleware.JWTWithConfig(middleware.JWTConfig{
		// 	SigningKey: []byte(handler.Key),
		// 	Skipper: func(c echo.Context) bool {
		// 		// Skip authentication for and signup login requests
		// 		if c.Path() == "/login" {
		// 			return true
		// 		}
		// 		return false
		// 	},
		// }),
	)

	// Initialize routes
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(s.APIPort))
}
