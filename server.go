package main

import (
	"github.com/atb/GolangMgoRestServer/server/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
