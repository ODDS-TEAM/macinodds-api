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
	// Authentication.
	e.POST("/login", a.Login)
	e.PATCH("/register", a.Register)
	e.POST("/logout", a.Logout)

	// Manage Devices.
	d := e.Group("/devices")
	d.POST("", a.CreateDevice)
	d.PUT("/:id", a.UpdateDevice)
	d.DELETE("/:id", a.RemoveDevice)
	d.GET("", a.GetDevices)
	d.GET("/:id", a.GetDevicesByID)
	// d.GET("/users/:uid", a.GetMyDevice)

	// Borrowings events.
	b := e.Group("/borrowings")
	b.GET("", a.GetBorrowings)
	b.GET("/users/:id", a.GetMyBorrowings)

	// Borrow and return Device.
	d.POST("/borrow/devices/:id", a.BorrowDevice)
	d.POST("/return/devices/:id", a.ReturnDevice)
	// d.POST("/:id/borrow", a.BorrowDevice)
	// d.POST("/:id/return", a.ReturnDevice)
}
