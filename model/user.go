package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User holds metadata about a user.
type (
	User struct {
		ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Role       string        `json:"role" bson:"role"`
		Name       string        `json:"name" bson:"name"`
		Email      string        `json:"email" bson:"email"`
		ImgProfile string        `json:"imgProfile,omitempty" bson:"imgProfile"`
		CreateAt   time.Time     `json:"createAt" bson:"createAt"`
		LastUpdate time.Time     `json:"lastUpdate" bson:"lastUpdate"`
	}
)

const (
	admin      = "admin"
	individual = "individual"
)

// IsAdmin return admin role of a user
func (u *User) IsAdmin() bool {
	return u.Role == admin
}

// GetName return name  of a user
func (u *User) GetName() string {
	return u.Name
}

// GetEmail return email of a user
func (u *User) GetEmail() string {
	return u.Email
}

// GetImgProfile image profile of a user
func (u *User) GetImgProfile() string {
	return u.ImgProfile
}
