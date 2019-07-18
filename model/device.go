package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Device holds metadata about a macbook.
	Device struct {
		ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		Serial     string        `json:"serial" bson:"serial"`
		Spec       string        `json:"spec" bson:"spec"`
		Img        string        `json:"img" bson:"img"`
		Location   string        `json:"location" bson:"location"`
		Borrowing  bool          `json:"borrowing" bson:"borrowing"`
		LastUpdate time.Time     `json:"lastUpdate" bson:"lastUpdate"`
	}

	DeviceList struct {
		ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name       string        `json:"name" bson:"name"`
		Serial     string        `json:"serial" bson:"serial"`
		Spec       string        `json:"spec" bson:"spec"`
		Img        string        `json:"img" bson:"img"`
		Location   string        `json:"location" bson:"location"`
		ReturnDate time.Time     `json:"returnDate" bson:"returnDate"`
		Borrowing  bool          `json:"borrowing" bson:"borrowing"`
		Borrower   Borrower2     `json:"borrower" bson:"borrower"`
	}

	Borrower2 struct {
		Name         string `json:"name" bson:"name"`
		Email        string `json:"email" bson:"email"`
		SlackAccount string `json:"slackAccount" bson:"slackAccount"`
		Tel          string `json:"tel" bson:"tel"`
	}
)
