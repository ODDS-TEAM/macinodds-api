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

API url : ``139.5.146.213:1323``

```
	e.GET("/", api.GetWelcome)
	e.POST("/signin", api.SignIn)

	m := e.Group("/mac")
	m.GET("", api.GetMac)
	m.GET("/:id", api.GetMacByID)
	m.POST("", api.CreateMac)
	m.PUT("/:id", api.UpdateMac)
	m.DELETE("/:id", api.RemoveMac)

	e.DELETE("/db/:id", api.RemoveMacInDB)
```