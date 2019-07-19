package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Borrowing holds metadata about a borrowing event.
	Borrowing struct {
		ID         bson.ObjectId `json:"_id" bson:"_id"`
		Date       time.Time     `json:"date" bson:"date"`
		Activity   string        `json:"activity" bson:"activity"`
		ReturnDate time.Time     `json:"returnDate" bson:"returnDate"`
		Memo       string        `json:"memo" bson:"memo"`
		Location   string        `json:"location" bson:"location"`
		Device     DeviceBorrow  `json:"device" bson:"device"`
		Borrower   Borrower      `json:"borrower" bson:"borrower"`
	}

	DeviceBorrow struct {
		ID   bson.ObjectId `json:"_id" bson:"_id"`
		Name string        `json:"name" bson:"name"`
	}

	Borrower struct {
		ID   bson.ObjectId `json:"_id" bson:"_id"`
		Name string        `json:"name" bson:"name"`
	}
)
