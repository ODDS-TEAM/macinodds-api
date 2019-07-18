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
	e.POST("/login", a.LogIn)
	e.PUT("/register", a.Register)
	e.POST("/logout", a.LogOut)

	d := e.Group("/devices")
	d.POST("", a.CreateDevice)
	d.PUT("/:id", a.UpdateDevice)
	d.DELETE("/:id", a.RemoveDevice)
	d.POST(":id/borrow", a.BorrowDevice)
	d.POST(":id/return", a.ReturnDevice)

	d.GET("", a.GetDevices)
	d.GET("/:id", a.GetDevicesByID)

	b := e.Group("/borrowings")
	b.GET("", a.GetBorrowings)
	b.GET("/users/:id", a.GetMyBorrowings)
}
