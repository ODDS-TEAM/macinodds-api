package api

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

// GetMac is
func (h *Handler) GetMac(c echo.Context) (err error) {
	m := []*model.Mac{}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("mac_odds_team").C("mac").Find(nil).Sort("-status", "-updateTime").All(&m); err != nil {
		return
	}

	return c.JSON(http.StatusOK, m)
}

// GetMacByID is a function of >>
func (h *Handler) GetMacByID(c echo.Context) (err error) {
	m := model.Mac{}
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
		return
	}

	return c.JSON(http.StatusOK, m)
}

// CreateMac is
func (h *Handler) CreateMac(c echo.Context) (err error) {
	// Create
	m := &model.Mac{
		ID:         bson.NewObjectId(),
		LastUpdate: time.Now(),
	}
	if err = c.Bind(m); err != nil {
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
	filePath := "/app/mac/" + imgName

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	m.Img = imgName

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	db := h.DB.Clone()
	defer db.Close()

	// Save mac in database
	if err = db.DB("mac_odds_team").C("mac").Insert(&m); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

// UpdateMac is
func (h *Handler) UpdateMac(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))
	nm := &model.Mac{
		ID:         id,
		LastUpdate: time.Now(),
	}

	if err = c.Bind(nm); err != nil {
		return
	}

	m := &model.Mac{}
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
		return
	}

	fmt.Println("new : ", nm.Img)
	fmt.Println("old : ", m.Img)

	if nm.Img != m.Img || nm.Img == "" {
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

		// Random filename, retaining existing extension.
		imgNewName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
		filePath := "/app/mac/" + imgNewName
		log.Println(filePath)

		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()

		nm.Img = imgNewName

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		imgName := m.Img
		if imgName != "" {
			fmt.Print("image remove empty")
			filePath = "/app/mac/" + imgName
			log.Println(filePath)

			// Remove image in Storage
			if err := os.Remove(filePath); err != nil {
				return err
			}
		}
	}

	// Validation
	if m.Name == "" || m.Serial == "" || m.Spec == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	// Update mac in database
	if err = db.DB("mac_odds_team").C("mac").Update(bson.M{"_id": id}, &nm); err != nil {
		return
	}

	return c.JSON(http.StatusOK, &nm)
}

// RemoveMac is
func (h *Handler) RemoveMac(c echo.Context) (err error) {
	m := model.Mac{}
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
		return
	}

	imgName := m.Img
	if imgName != "" {

		filePath := "/app/mac/" + imgName
		log.Println(filePath)

		// Remove image in Storage
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	// Remove mac in DB
	if err = db.DB("mac_odds_team").C("mac").RemoveId(id); err != nil {
		return
	}

	return c.JSON(http.StatusOK, err)
}
