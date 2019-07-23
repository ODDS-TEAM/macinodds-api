package api

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

type ReturnDate struct {
	ReturnDate string `json:"returnDate" bson:"returnDate"`
}

func (db *MongoDB) BorrowDevice(c echo.Context) (err error) {
	id := getID(c)
	q := bson.M{
		"_id": id,
	}
	ch := bson.M{
		"$set": bson.M{
			"borrowing": true,
		},
	}

	if err := db.DCol.Update(q, &ch); err != nil {
		return err
	}
	m := &model.Device{}
	if err := db.DCol.FindId(id).One(&m); err != nil {
		return err
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["id"].(string)
	qid := bson.ObjectIdHex(uid)

	u := &model.User{}
	if err := db.UCol.FindId(qid).One(&u); err != nil {
		return err
	}

	log.Println("uid", uid)
	log.Println(&u)

	var r ReturnDate
	if err := c.Bind(&r); err != nil {
		return err
	}
	log.Print("Sivaroot R", r)
	t, err := time.Parse(time.RFC3339, r.ReturnDate)

	b := &model.Borrowing{
		ID:         bson.NewObjectId(),
		Date:       time.Now(),
		Activity:   "borrow",
		ReturnDate: t,  // <<<<
		Memo:       "", // <<<
		Location:   "", // <<<
		Device: model.Name{
			ID: m.ID,
		},
		Borrower: model.Name{
			ID: u.ID,
		},
	}

	if err := db.BCol.Insert(b); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, b)
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
