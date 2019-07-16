package route

import (
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
	d := e.Group("/devices")
	d.POST("", a.CreateDevice)
	d.PUT("/:id", a.UpdateDevice)
	d.DELETE("/:id", a.RemoveDevice)
	d.GET("", a.GetDevices)

	// b := e.Group("/borrowings")
	// b.GET("", a.GetBorrowings)

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
	// m.GET("/:id", api.GetMacsByID)
	// m.GET("/users/:uid", api.GetMacsByUID)

	// m.POST("/users/:uid/borrow", api.BorrowMac)
	// m.POST("/users/:uid/return", api.ReturnMac)

	// m.GET("/:id", api.GetMacByID)

	// m.POST("", api.CreateMac)
	// m.PUT("/:id", api.UpdateMac)
	// m.DELETE("/:id", api.RemoveMac)
}
