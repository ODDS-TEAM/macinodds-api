package model

import "gopkg.in/mgo.v2/bson"

type (
	// Borrowing is
	Borrowing struct {
		ID         bson.ObjectId `json:"_id"`
		Date       string        `json:"date" bson:"date"`
		Activity   string        `json:"activity" bson:"activity"`
		UserID     bson.ObjectId `json:"userID bson:"userID"`
		MacID      bson.ObjectId `json:"macID" bson:"macID"`
		ReturnDate string        `json:"returnDate" bson:"returnDate"`
		Memo       string        `json:"memo" bson:"memo"`
		Location   string        `json:"location" bson:"location"`
	}
)
