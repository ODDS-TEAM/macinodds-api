package route

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.odds.team/internship/macinodds-api/api"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	db, err := api.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &api.MongoDB{
		Conn: db.Conn,
		DCol: db.DCol,
		UCol: db.UCol,
		BCol: db.BCol,
	}

	// Routes
	// Respond to API health checks.
	// Indicate the server is healthy.
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "mac.odds.team : ok!")
	})

	// LOGIN
	// e.POST("/loginGoogle", api.LoginGoogle)
	// e.POST("/register", api.Register)
	// e.POST("/login", api.Login)
	// e.POST("/logout", api.Logout)

	// HISTORY
	// h := e.Group("/histories")
	// h.GET("", api.GetHistories)
	// h.GET("/users/:uid", api.GetHistoriesByUID)

	// MAC
	m := e.Group("/devices")
	m.GET("", a.GetDevices)

	// m.GET("/:id", api.GetMacsByID)
	// m.GET("/users/:uid", api.GetMacsByUID)

	m.POST("", a.CreateDevice)
	// m.PUT("/:id", a.UpdateDevice)
	m.DELETE("/:id", a.RemoveDevice)
	// m.POST("/users/:uid/borrow", api.BorrowMac)
	// m.POST("/users/:uid/return", api.ReturnMac)

	// m.GET("/:id", api.GetMacByID)

	// m.POST("", api.CreateMac)
	// m.PUT("/:id", api.UpdateMac)
	// m.DELETE("/:id", api.RemoveMac)
}
