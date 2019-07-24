package api

import (
	"time"

	model "gitlab.odds.team/internship/macinodds-api/model"
	"google.golang.org/api/oauth2/v1"
	"gopkg.in/mgo.v2/bson"
)

// FindUser search and validation of user data.
func (db *MongoDB) FindUser(tokeninfo *oauth2.Tokeninfo) (*model.User, string) {
	u := &model.User{}
	status := "new"
	q := bson.M{
		"email": tokeninfo.Email,
	}
	if err := db.UCol.Find(q).One(&u); err != nil {
		return nil, status
	}

	if u.Name == "" || u.SlackAccount == "" || u.Tel == "" || u.ImgProfile == "" {
		status = "notComplete"
	} else {
		status = "exist"
	}

	return u, status
}

// CreateUser created by ID and email in database.
func (db *MongoDB) CreateUser(tokeninfo *oauth2.Tokeninfo) *model.User {
	u := &model.User{
		ID:       bson.NewObjectId(),
		Role:     "individual", // default user's role
		Email:    tokeninfo.Email,
		CreateAt: time.Now(),
	}
	if err := db.UCol.Insert(&u); err != nil {
		return nil
	}

	return u
}

// UpdateUser or register information of user to database.
func (db *MongoDB) UpdateUser(uid bson.ObjectId, u *model.User) {
	q := bson.M{
		"_id":   uid,
		"email": u.Email,
	}
	ch := bson.M{
		"$set": bson.M{
			"name":         u.Name,
			"imgProfile":   u.ImgProfile,
			"slackAccount": u.SlackAccount,
			"tel":          u.Tel,
		},
	}

	if err := db.UCol.Update(q, &ch); err != nil {
		return
	}
}
