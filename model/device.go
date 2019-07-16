package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Device holds metadata about a macbook.
type (
	Device struct {
		ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		Serial     string        `json:"serial" bson:"serial"`
		Spec       string        `json:"spec" bson:"spec"`
		Img        string        `json:"img" bson:"img"`
		Location   string        `json:"location" bson:"location"`
		LastUpdate time.Time     `json:"lastUpdate" bson:"lastUpdate"`
		Borrowing  bool          `json:"borrowing" bson:"borrowing"`
	}
)
