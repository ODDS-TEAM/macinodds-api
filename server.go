package main

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/devices-api/route"
)

func main() {
	e := echo.New()

	// Connects to the database

	// Initial routes
	route.Init(e)

	// Runs the server
	e.Logger.Fatal(e.Start(":1323"))
}
