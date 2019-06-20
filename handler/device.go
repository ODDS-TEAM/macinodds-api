package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/models"
	"gopkg.in/mgo.v2/bson"
)

type (
	form struct {
		ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name  string        `json:"name" bson:"name"`
		Email string        `json:"email" bson:"email"`
		File  string        `json:"file" bson:"file"`
		Size  int64         `json:"size"`
	}
)

// Handler
func (h *Handler) GetAPI(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Welcome to Macinodds devices API v.1.0!")
}

func (h *Handler) GetDevices(c echo.Context) (err error) {
	dv := []*model.Device{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{}).All(&dv); err != nil {
		return
	}

	return c.JSON(http.StatusOK, dv)
}

func (h *Handler) GetByID(c echo.Context) (err error) {
	dv := model.Device{}
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&dv); err != nil {
		return
	}

	return c.JSON(http.StatusOK, dv)
}

func (h *Handler) CreateDevice(c echo.Context) (err error) {
	// Source
	file, err := c.FormFile("img")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// random filename, retaining existing extension.
	imgName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
	log.Println(imgName)
	filePath := "/app/devices/" + imgName
	// filePath := imgName

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	dv := &model.Device{
		ID:  bson.NewObjectId(),
		Img: imgName,
	}
	if err = c.Bind(dv); err != nil {
		return err
	}

	// Validation
	if dv.Name == "" || dv.Serial == "" || dv.Spec == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	db := h.DB.Clone()
	defer db.Close()

	// Save device in database
	if err = db.DB("macinodds").C("devices").Insert(dv); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dv)
}

func (h *Handler) UpdateDevice(c echo.Context) (err error) {
	dv := &model.Device{}
	id := bson.ObjectIdHex(c.Param("id"))

	if err = c.Bind(dv); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()

	// Validation
	if dv.Name == "" || dv.Serial == "" || dv.Spec == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	// Update device in database
	if err = db.DB("macinodds").C("devices").Update(bson.M{"_id": id}, &dv); err != nil {
		return
	}

	dv.ID = id

	return c.JSON(http.StatusOK, dv)
}

func (h *Handler) RemoveDevice(c echo.Context) (err error) {
	db := h.DB.Clone()
	id := bson.ObjectIdHex(c.Param("id"))
	defer db.Close()

	if err = db.DB("macinodds").C("devices").RemoveId(id); err != nil {
		return
	}

	return err
}
