package api

import "gopkg.in/mgo.v2"

type (
	// Handler to communication session with the database.
	Handler struct {
		DB *mgo.Session
	}
)
