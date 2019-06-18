package model

import "gopkg.in/mgo.v2/bson"

// Device structural
type (
	Device struct {
		ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name   string        `json:"name" bson:"name"`
		Serial string        `json:"serial" bson:"serial"`
	}
)
