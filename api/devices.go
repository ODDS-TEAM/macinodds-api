package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

// GetDevices show a list of all Devices and sorted by borrowing status and last update time.
func (db *MongoDB) GetDevices(c echo.Context) (err error) {
	m, err := db.GetDevicesDB()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &m)
}

// GetDevicesDB is
func (db *MongoDB) GetDevicesDB() ([]*model.Device, error) { //ListBooks
	var m []*model.Device
	if err := db.DCol.Find(nil).Sort("borrowing", "-lastUpdate").All(&m); err != nil {
		return nil, err
	}

	return m, nil
}

// CreateDevice create a Mac into the database.
func (db *MongoDB) CreateDevice(c echo.Context) (err error) {
	m := &model.Device{
		ID:         bson.NewObjectId(),
		LastUpdate: time.Now(),
	}
	if err = c.Bind(m); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, m)

	// Validate

	// Source
	// file, err := c.FormFile("img")
	// if err != nil {
	// 	return &echo.HTTPError{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "invalid to or message fields",
	// 	}
	// }

	// src, err := file.Open()
	// if err != nil {
	// 	return err
	// }
	// defer src.Close()

	// random filename, retaining existing extension.
	// imgName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
	// log.Println(imgName)
	// filePath := "app/mac/" + imgName

	// dst, err := os.Create(filePath)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// m.Img = imgName

	// Copy
	// if _, err = io.Copy(dst, src); err != nil {
	// 	return err
	// }

	// db := h.DB.Clone()
	// defer db.Close()

	// Save Mac in database
	// if err = db.DB("mac_odds_team").C("mac").Insert(&m); err != nil {
	// 	return err
	// }

	// return c.JSON(http.StatusCreated, m)
}

// GetMacByID show the Mac by ID.
// func (h *HandlerDB) GetMacByID(c echo.Context) (err error) {
// 	id := bson.ObjectIdHex(c.Param("id"))
// 	m := model.Mac{}

// 	db := h.DB.Clone()
// 	defer db.Close()

// 	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
// 		return
// 	}

// 	return c.JSON(http.StatusOK, m)
// }

// UpdateMac update Mac that has been modified.
// func (h *HandlerDB) UpdateMac(c echo.Context) (err error) {
// 	id := bson.ObjectIdHex(c.Param("id"))
// 	nm := &model.Mac{
// 		ID:         id,
// 		LastUpdate: time.Now(),
// 	}

// 	if err = c.Bind(nm); err != nil {
// 		return
// 	}

// 	m := &model.Mac{}

// 	db := h.DB.Clone()
// 	defer db.Close()

// 	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
// 		return
// 	}

// 	fmt.Println("new : ", nm.Img)
// 	fmt.Println("old : ", m.Img)

// 	if nm.Img != m.Img || nm.Img == "" {
// 		fmt.Println("Image Not Math")
// 		// Source
// 		file, err := c.FormFile("img")
// 		if err != nil {
// 			return &echo.HTTPError{
// 				Code:    http.StatusBadRequest,
// 				Message: "invalid to or message fields",
// 			}
// 		}

// 		src, err := file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer src.Close()

// 		// Random filename, retaining existing extension.
// 		imgNewName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
// 		filePath := "app/mac/" + imgNewName
// 		log.Println(filePath)

// 		dst, err := os.Create(filePath)
// 		if err != nil {
// 			return err
// 		}
// 		defer dst.Close()

// 		nm.Img = imgNewName

// 		// Copy
// 		if _, err = io.Copy(dst, src); err != nil {
// 			return err
// 		}

// 		imgName := m.Img
// 		if imgName != "" {
// 			fmt.Print("image remove empty")
// 			filePath = "app/mac/" + imgName
// 			log.Println(filePath)

// 			// Remove image in Storage
// 			if err := os.Remove(filePath); err != nil {
// 				return err
// 			}
// 		}
// 	}

// Validation
// 	if m.Name == "" || m.Serial == "" || m.Spec == "" {
// 		return &echo.HTTPError{
// 			Code:    http.StatusBadRequest,
// 			Message: "invalid to or message fields",
// 		}
// 	}

// 	// Update Mac in database
// 	if err = db.DB("mac_odds_team").C("mac").Update(bson.M{"_id": id}, &nm); err != nil {
// 		return
// 	}

// 	return c.JSON(http.StatusOK, &nm)
// }

// RemoveMac remove the Mac's data.
// func (h *HandlerDB) RemoveMac(c echo.Context) (err error) {
// 	id := bson.ObjectIdHex(c.Param("id"))
// 	m := model.Mac{}

// 	db := h.DB.Clone()
// 	defer db.Close()

// 	if err = db.DB("mac_odds_team").C("mac").Find(bson.M{"_id": id}).One(&m); err != nil {
// 		return
// 	}

// 	// Remove image in Storage
// 	imgName := m.Img
// 	if imgName != "" {
// 		filePath := "app/mac/" + imgName

// 		if err = os.Remove(filePath); err != nil {
// 			return
// 		}
// 	}

// 	// Remove Mac in database
// 	if err = db.DB("mac_odds_team").C("mac").RemoveId(id); err != nil {
// 		return
// 	}

// 	return c.JSON(http.StatusOK, "the mac deleted successfully")
// }
