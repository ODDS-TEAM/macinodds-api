package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Device holds metadata about a macbook.
	Device struct {
		ID         bson.ObjectId `json:"_id,omitempty" bson:"_id"`
		Name       string        `json:"name" bson:"name"`
		Serial     string        `json:"serial" bson:"serial"`
		Spec       string        `json:"spec" bson:"spec"`
		Img        string        `json:"img" bson:"img"`
		Location   string        `json:"location" bson:"location"`
		Borrowing  bool          `json:"borrowing" bson:"borrowing"`
		LastUpdate time.Time     `json:"lastUpdate" bson:"lastUpdate"`
		ReturnDate time.Time     `json:"returnDate" bson:"returnDate"`
		Borrower   Borrower      `json:"borrower" bson:"borrower"`
	}

	// Borrower holds information about a borrower.
	Borrower struct {
		ID           bson.ObjectId `json:"_id,omitempty" bson:"_id"`
		Name         string        `json:"name,omitempty" bson:"name,omitempty"`
		Email        string        `json:"email,omitempty" bson:"email,omitempty"`
		SlackAccount string        `json:"slackAccount,omitempty" bson:"slackAccount,omitempty"`
		Tel          string        `json:"tel,omitempty" bson:"tel,omitempty"`
	}
)
