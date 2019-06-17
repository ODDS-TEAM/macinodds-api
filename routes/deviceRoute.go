package deviceRoute

import "github.com/labstack/echo"

func Init(e *echo.Echo) {
	e.GET("/api/devices", deviceController.GetDevices)
	// e.GET("/api/device/:id", deviceController.GetById)
	// e.POST("/api/device", deviceController.NewTodo)
	// e.DELETE("/api/device/:id", deviceController.RemoveTodo)
}
