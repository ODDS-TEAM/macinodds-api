# Install Golang : OS X

##### Install brew

Install brew from [brew.sh](http://brew.sh/)

##### Install Git

```
$ brew install git 
```

##### Install Go

Install golang binaries using `brew`

```
$ brew install go
```

##### Testing it all

```
$ go env
$ go version
```

##### Pull package

```
$ go get
```

##### Run

```
$ go run server.go
```

##### API routes

API url : `139.5.146.213:1323`

```
	e.GET("/api", h.GetAPI)
	e.GET("/api/devices", h.GetDevices)
	e.GET("/api/devices/:id", handler.GetByID)
	e.POST("/api/devices", h.CreateDevice)
	e.PUT("/api/devices/:id", handler.EditDevice)
	e.DELETE("/api/divices/:id", handler.RemoveDevice)
```