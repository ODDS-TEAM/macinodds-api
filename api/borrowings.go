package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

// ReturnDate is
type ReturnDate struct {
	ReturnDate string `json:"returnDate" bson:"returnDate"`
}

// BorrowDevice is
func (db *MongoDB) BorrowDevice(c echo.Context) (err error) {
	// edit status device
	uid := GetIDFromToken(c)
	id := getID(c)

	// check devices false and not have borrower
	quid := bson.M{
		"borrower._id": uid,
	}
	qid := bson.M{
		"_id":       id,
		"borrowing": false,
	}

	count, _ := db.DCol.Find(quid).Count()
	if count > 0 {
		return c.JSON(http.StatusNotFound, "You can not borrow device more than one unit")
	}
	d := &model.Device{}
	if err = db.DCol.Find(qid).One(d); err != nil {
		return c.JSON(http.StatusNotFound, "Can not find the device or device not wang")
	}

	r := &ReturnDate{}
	if err := c.Bind(r); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, r.ReturnDate)
	if err != nil {
		return err
	}

	// update data device
	q := bson.M{
		"_id": id,
	}
	ch := bson.M{
		"$set": bson.M{
			"borrowing": true,
			"borrower": bson.M{
				"_id": uid,
			},
			"returnDate": t,
		},
	}
	if err := db.DCol.Update(q, &ch); err != nil {
		return err
	}

	b := &model.Borrowing{
		ID:         bson.NewObjectId(),
		Date:       time.Now(),
		Activity:   "borrow",
		ReturnDate: t,  // <<<<
		Memo:       "", // <<<
		Location:   "", // <<<
		Device: model.Name{
			ID: id,
		},
		Borrower: model.Name{
			ID: uid,
		},
	}

	if err := db.BCol.Insert(b); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "yeh")
}

func (db *MongoDB) ReturnDevice(c echo.Context) (err error) {
	id := getID(c)
	q := bson.M{
		"_id": id,
	}
	ch := bson.M{
		"$set": bson.M{
			"borrowing": false,
		},
	}

	if err := db.DCol.Update(q, &ch); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Return "+id)
}

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
