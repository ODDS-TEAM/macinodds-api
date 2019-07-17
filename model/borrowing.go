package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Borrowing holds metadata about a borrowing event.
	Borrowing struct {
		// ID         bson.ObjectId `json:"_id" bson:"_id"`
		// Date       time.Time     `json:"date" bson:"date"`
		Activity   string        `json:"activity" bson:"activity"`
		MacID      bson.ObjectId `json:"macID" bson:"macID"`
		ReturnDate time.Time     `json:"returnDate" bson:"returnDate"`
		Memo       string        `json:"memo" bson:"memo"`
		Location   string        `json:"location" bson:"location"`
		// Borrower   struct {
		// 	ID           bson.ObjectId `json:"_id" bson:"_id"`
		// 	Name         string        `json:"name" bson:"name"`
		// 	SlackAccount string        `json:"slackAccount" bson:"slackAccount"`
		// 	Tel          string        `json:"tel" bson:"tel"`
		// }
	}
)
