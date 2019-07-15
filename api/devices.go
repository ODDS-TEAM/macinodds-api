package api

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

// CreateDevice create a device into the database.
func (db *MongoDB) CreateDevice(c echo.Context) (err error) {
	m, err := db.InsertDevice(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

// GetDevices show a list of all Devices and sorted by borrowing status and last update time.
func (db *MongoDB) GetDevices(c echo.Context) (err error) {
	m, err := db.FindDevices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &m)
}

func (db *MongoDB) FindDevices() ([]*model.Device, error) { //ListBooks
	m := []*model.Device{}
	if err := db.DCol.Find(nil).Sort("borrowing", "-lastUpdate").All(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func GenImgID(f string) (string, string) {
	i := uuid.Must(uuid.NewV4()).String() + path.Ext(f)
	p := "/app/mac/" + i

	return i, p
}

func OpenFile(c echo.Context) (*multipart.FileHeader, multipart.File) {
	f, err := c.FormFile("img")
	if err != nil {
		return nil, nil
	}
	s, err := f.Open()
	if err != nil {
		return nil, nil
	}
	defer s.Close()

	return f, s
}

func CreateFile(p string, s multipart.File) {
	d, err := os.Create(p)
	if err != nil {
		return
	}
	defer d.Close()
	if _, err = io.Copy(d, s); err != nil {
		return
	}
}

func (db *MongoDB) InsertDevice(c echo.Context) (*model.Device, error) { //ListBooks
	file, src := OpenFile(c)
	imgName, filePath := GenImgID(file.Filename)
	CreateFile(filePath, src)

	// create
	m := &model.Device{
		ID:         bson.NewObjectId(),
		Img:        imgName,
		LastUpdate: time.Now(),
		Borrowing:  false,
	}
	if err := c.Bind(m); err != nil {
		return nil, err
	}

	// Save the device in database
	if err := db.DCol.Insert(&m); err != nil {
		return nil, err
	}

	return m, nil
}

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
