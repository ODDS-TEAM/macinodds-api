package api

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MongoDB holds metadata about session database and collections name.
type (
	MongoDB struct {
		Conn *mgo.Session
		DCol *mgo.Collection
		UCol *mgo.Collection
		BCol *mgo.Collection
	}
)

// NewMongoDB creates a new macOddsTeamDB backed by a given Mongo server.
func NewMongoDB() (*MongoDB, error) {
	conn, err := mgo.Dial("mac.odds.team:27017")

	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	return &MongoDB{
		Conn: conn,
		DCol: conn.DB("macOddsTeamDB").C("devices"),
		UCol: conn.DB("macOddsTeamDB").C("users"),
		BCol: conn.DB("macOddsTeamDB").C("borrowings"),
	}, nil
}

// Close closes the database.
func (db *MongoDB) Close() {
	db.Conn.Close()
}
