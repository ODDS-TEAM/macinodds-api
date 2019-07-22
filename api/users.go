package api

import (
	"time"

	model "gitlab.odds.team/internship/macinodds-api/model"
	"google.golang.org/api/oauth2/v1"
	"gopkg.in/mgo.v2/bson"
)

// FindUser is a function to finding the user by id
func (db *MongoDB) FindUser(tokeninfo *oauth2.Tokeninfo) *model.User {
	u := &model.User{}
	q := bson.M{
		"email": tokeninfo.Email,
	}
	if err := db.UCol.Find(q).One(&u); err != nil {
		return nil
	}

	return u
}

// CreateUser in database.
func (db *MongoDB) CreateUser(tokeninfo *oauth2.Tokeninfo) *model.User {
	u := &model.User{
		ID:       bson.NewObjectId(),
		Role:     "individual",
		Email:    tokeninfo.Email,
		CreateAt: time.Now(),
	}
	if err := db.UCol.Insert(&u); err != nil {
		return nil
	}

	return u
}

// UpdateUser to database
func (db *MongoDB) UpdateUser(q bson.M, ch bson.M) {
	if err := db.UCol.Update(q, &ch); err != nil {
		return
	}
}
