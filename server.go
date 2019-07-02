package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	route "gitlab.odds.team/internship/macinodds-api/routes"
)

func main() {
	// Use labstack/echo for rich routing.
	// See https://echo.labstack.com/
	e := echo.New()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
		// middleware.JWTWithConfig(middleware.JWTConfig{
		// 	SigningKey: []byte(handler.Key),
		// 	Skipper: func(c echo.Context) bool {
		// 		// Skip authentication for and signin requests
		// 		if c.Path() == "/signin" {
		// 			return true
		// 		}
		// 		return false
		// 	},
		// }),
	)

	// Initialize routes
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
