package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	route "gitlab.odds.team/internship/macinodds-api/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
