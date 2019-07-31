package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User holds metadata about a user.
type (
	User struct {
		ID           bson.ObjectId `json:"_id,omitempty" bson:"_id"`
		Role         string        `json:"role" bson:"role"`
		Name         string        `json:"name" bson:"name"`
		Email        string        `json:"email" bson:"email"`
		ImgProfile   string        `json:"imgProfile" bson:"imgProfile"`
		SlackAccount string        `json:"slackAccount" bson:"slackAccount"`
		Tel          string        `json:"tel" bson:"tel"`
		CreateAt     time.Time     `json:"createAt,omitempty" bson:"createAt"`
	}
)
