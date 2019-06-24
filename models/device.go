package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Device holds metadata about a device.
type (
	Device struct {
		ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		Serial     string        `json:"serial" bson:"serial"`
		Spec       string        `json:"spec" bson:"spec"`
		Status     bool          `json:"status" bson:"status"`
		Holder     string        `json:"holder" bson:"holder"`
		Tel        string        `json:"tel" bson:"tel"`
		Img        string        `json:"img" bson:"img"`
		UpdateTime time.Time     `json:"updateTime" bson:"updateTime"`
	}
)
