package api

import (
	"log"
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

	return c.JSON(http.StatusOK, "Borrow!")
}

type (
	// ReturnData is
	ReturnData struct {
		Memo     string `json:"memo" bson:"memo"`
		Location string `json:"location" bson:"location"`
	}
)

// ReturnDevice is
func (db *MongoDB) ReturnDevice(c echo.Context) (err error) {
	uid := GetIDFromToken(c)
	id := getID(c)
	// check mathcing user and device
	q := bson.M{
		"borrower._id": uid,
		"_id":          id,
		"borrowing":    true,
	}
	d := &model.Device{}
	if err = db.DCol.Find(q).One(d); err != nil {
		return c.JSON(http.StatusNotFound, "Not your device")
	}

	r := &ReturnData{}
	if err := c.Bind(r); err != nil {
		return err
	}
	defer log.Println("res", r)

	// update data device
	q2 := bson.M{
		"_id": id,
	}
	ch := bson.M{
		"$set": bson.M{
			"location":   r.Location,
			"borrowing":  false,
			"borrower":   bson.M{},
			"returnDate": time.Time{},
		},
	}
	if err := db.DCol.Update(q2, &ch); err != nil {
		return err
	}
	// update borrowing database
	b := &model.Borrowing{
		ID:         bson.NewObjectId(),
		Date:       time.Now(),
		Activity:   "return",
		ReturnDate: time.Now(), // <<<<
		Memo:       r.Memo,     // <<<
		Location:   r.Location, // <<<
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

	return c.JSON(http.StatusOK, "Return!")
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
	b := []*model.Borrowing{}
	dLookup := bson.M{
		"$lookup": bson.M{
			"from":         "devices",
			"localField":   "device._id",
			"foreignField": "_id",
			"as":           "device",
		},
	}
	uLookup := bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "borrower._id",
			"foreignField": "_id",
			"as":           "borrower",
		},
	}
	dUnwind := bson.M{
		"$unwind": bson.M{
			"path":                       "$device",
			"preserveNullAndEmptyArrays": true,
		},
	}
	uUnwind := bson.M{
		"$unwind": bson.M{
			"path":                       "$borrower",
			"preserveNullAndEmptyArrays": true,
		},
	}

	q := []bson.M{dLookup, uLookup, dUnwind, uUnwind}
	e, _ := id.MarshalText()
	if len(e) > 0 {
		match := bson.M{
			"$match": bson.M{
				"borrower._id": id,
			},
		}
		q = []bson.M{match, dLookup, uLookup, dUnwind, uUnwind}
	}

	if err := db.BCol.Pipe(q).All(&b); err != nil {
		return nil
	}

	return b
}
