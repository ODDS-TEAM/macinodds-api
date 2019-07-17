package api

import (
	"net/http"

	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

func (db *MongoDB) GetBorrowings(c echo.Context) (err error) {
	b := db.findBorrowingsDB(c, "")

	return c.JSON(http.StatusCreated, b)
}

func (db *MongoDB) GetMyBorrowings(c echo.Context) (err error) {
	id := getID(c)
	b := db.findBorrowingsDB(c, id)
	return c.JSON(http.StatusCreated, b)
}

func (db *MongoDB) findBorrowingsDB(c echo.Context, id bson.ObjectId) []*model.Borrowing {
	var q bson.M
	e, _ := id.MarshalText()

	if len(e) > 0 {
		q = bson.M{"borrower._id": id}
	} else {
		q = nil
	}

	b := []*model.Borrowing{}

	if err := db.BCol.Find(q).Sort("-date").All(&b); err != nil {
		return nil
	}

	return b
}
