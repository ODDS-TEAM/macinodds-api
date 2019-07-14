# API mac.odds.team

## Install Golang : OS X

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
$ go version (1.12.6)
```

##### Pull package

```
$ go get
```

##### Run

```
$ go run server.go
```

##### Build & Run

```
$ go build server.go
$ ./server
```

##### API routes

API url : [mac.odds.team/api](http://mac.odds.team/api)

```
	e.GET("/", api.GetWelcome)

	m := e.Group("/mac")
	m.GET("", api.GetMac)
	m.GET("/:id", api.GetMacByID)
	
	m.POST("", api.CreateMac)
	m.PUT("/:id", api.UpdateMac)
	m.DELETE("/:id", api.RemoveMac)
```