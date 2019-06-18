package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/models"
	"gopkg.in/mgo.v2/bson"
)

// Handler
func (h *Handler) GetAPI(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Welcome to Macinodds devices API v.1.0!")
}

func (h *Handler) GetDevices(c echo.Context) (err error) {
	devices := []*model.Device{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{}).All(&devices); err != nil {
		return
	}

	return c.JSON(http.StatusOK, devices)
}

func (h *Handler) GetByID(c echo.Context) (err error) {
	device := model.Device{}
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&device); err != nil {
		return
	}

	return c.JSON(http.StatusOK, device)
}

func (h *Handler) CreateDevice(c echo.Context) (err error) {
	dv := &model.Device{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(dv); err != nil {
		return
	}

	// Validation
	if dv.Name == "" || dv.Serial == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	db := h.DB.Clone()
	defer db.Close()

	// Save post in database
	if err = db.DB("macinodds").C("devices").Insert(dv); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, dv)
}

func (h *Handler) RemoveDevice(c echo.Context) (err error) {
	db := h.DB.Clone()
	id := c.Param("id")
	idoi := bson.ObjectIdHex(id)
	defer db.Close()

	fmt.Println("todo ID = " + idoi)

	if err = db.DB("macinodds").C("devices").RemoveId(idoi); err != nil {
		return
	}

	return err
}
