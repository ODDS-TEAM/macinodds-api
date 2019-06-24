package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/models"
	"gopkg.in/mgo.v2/bson"
)

// GetAPI is
func (h *Handler) GetAPI(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Welcome to Macinodds devices API v.1.0!")
}

// GetDevices is
func (h *Handler) GetDevices(c echo.Context) (err error) {
	dv := []*model.Device{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(nil).Sort("-status", "-updateTime").All(&dv); err != nil {
		return
	}

	return c.JSON(http.StatusOK, dv)
}

// GetByID is a function of >>
func (h *Handler) GetByID(c echo.Context) (err error) {
	dv := model.Device{}
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{"_id": id}).One(&dv); err != nil {
		return
	}

	return c.JSON(http.StatusOK, dv)
}

// CreateDevice is
func (h *Handler) CreateDevice(c echo.Context) (err error) {
	// Create
	dv := &model.Device{
		ID:         bson.NewObjectId(),
		UpdateTime: time.Now(),
	}
	if err = c.Bind(dv); err != nil {
		return err
	}

	// Validate

	// Source
	file, err := c.FormFile("img")
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
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

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	dv.Img = imgName

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	db := h.DB.Clone()
	defer db.Close()

	// Save device in database
	if err = db.DB("macinodds").C("devices").Insert(&dv); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dv)
}

// UpdateDevice is
func (h *Handler) UpdateDevice(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))
	ndv := &model.Device{
		ID:         id,
		UpdateTime: time.Now(),
	}

	if err = c.Bind(ndv); err != nil {
		return
	}

	dv := &model.Device{}
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{"_id": id}).One(&dv); err != nil {
		return
	}

	fmt.Println("new : ", ndv)
	fmt.Println("old : ", dv)
	fmt.Println("new : ", ndv.Img)
	fmt.Println("old : ", dv.Img)

	if ndv.Img != dv.Img || ndv.Img == "" {
		fmt.Println("Image Not Math")
		// Source
		file, err := c.FormFile("img")
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "invalid to or message fields",
			}
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// random filename, retaining existing extension.
		imgNewName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
		filePath := "/app/devices/" + imgNewName
		log.Println(filePath)

		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()

		ndv.Img = imgNewName

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		imgName := dv.Img
		if imgName != "" {
			fmt.Print("image remove empty")
			filePath = "/app/devices/" + imgName
			log.Println(filePath)

			// Remove image in Storage
			if err := os.Remove(filePath); err != nil {
				return err
			}
		}
	}

	// Validation
	if dv.Name == "" || dv.Serial == "" || dv.Spec == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	// Update device in database
	if err = db.DB("macinodds").C("devices").Update(bson.M{"_id": id}, &ndv); err != nil {
		return
	}

	return c.JSON(http.StatusOK, &ndv)
}

// RemoveDevice is
func (h *Handler) RemoveDevice(c echo.Context) (err error) {
	dv := model.Device{}
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("macinodds").C("devices").Find(bson.M{"_id": id}).One(&dv); err != nil {
		return
	}

	imgName := dv.Img
	if imgName != "" {

		filePath := "/app/devices/" + imgName
		log.Println(filePath)

		// Remove image in Storage
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	// Remove device in DB
	if err = db.DB("macinodds").C("devices").RemoveId(id); err != nil {
		return
	}

	return c.JSON(http.StatusOK, err)
}

// RemoveDB is
func (h *Handler) RemoveDB(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	// Remove device in DB
	if err = db.DB("macinodds").C("devices").RemoveId(id); err != nil {
		return
	}

	return c.JSON(http.StatusOK, err)
}
