package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Borrowing holds metadata about a borrowing event.
	Borrowing struct {
		ID         bson.ObjectId `json:"_id,omitempty" bson:"_id"`
		Date       time.Time     `json:"date" bson:"date"`
		Activity   string        `json:"activity" bson:"activity"`
		ReturnDate time.Time     `json:"returnDate" bson:"returnDate"`
		Memo       string        `json:"memo" bson:"memo"`
		Location   string        `json:"location" bson:"location"`
		Device     Name          `json:"device" bson:"device"`
		Borrower   Name          `json:"borrower" bson:"borrower"`
	}

	// Name of device or borrower when query.
	Name struct {
		ID   bson.ObjectId `json:"_id,omitempty" bson:"_id"`
		Name string        `json:"name,omitempty" bson:"name,omitempty"`
	}
)
