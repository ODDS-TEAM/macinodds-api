package api

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2/bson"
)

// CreateDevice create a device into the database.
func (db *MongoDB) CreateDevice(c echo.Context) (err error) {
	m, err := db.insertDeviceDB(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

// UpdateDevice updates the entry for a given device.
func (db *MongoDB) UpdateDevice(c echo.Context) (err error) {
	id := getID(c)
	nm := &model.Device{
		ID:         id,
		LastUpdate: time.Now(),
	}
	if err = c.Bind(nm); err != nil {
		return
	}

	m := &model.Device{}
	if err := db.DCol.Find(bson.M{"_id": id}).One(&m); err != nil {
		return err
	}

	log.Println("new:", nm.Img)
	log.Println("old:", m.Img)

	if nm.Img != m.Img || nm.Img == "" {
		// Source
		file, src, _ := openFile(c)

		// Random filename, retaining existing extension.
		imgNewName, filePath := genImgID(file.Filename)

		createFile(filePath, src)

		nm.Img = imgNewName

		if m.Img != "" {
			removeFile(m)
		}
	}

	// Update Mac in database
	db.updateDeviceDB(id, nm)

	return c.JSON(http.StatusOK, &nm)
}

// RemoveDevice removes a given device by its ID.
func (db *MongoDB) RemoveDevice(c echo.Context) (err error) {
	db.removeDeviceDB(c)

	return c.JSON(http.StatusOK, "The device deleted successfully")
}

// GetDevices show a list of all Devices and sorted by borrowing status and last update time.
func (db *MongoDB) GetDevices(c echo.Context) (err error) {
	m, err := db.findDevicesDB()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &m)
}

func (db *MongoDB) insertDeviceDB(c echo.Context) (*model.Device, error) {
	file, src, _ := openFile(c)
	imgName, filePath := genImgID(file.Filename)
	createFile(filePath, src)

	m := &model.Device{
		ID:         bson.NewObjectId(),
		Img:        imgName,
		LastUpdate: time.Now(),
		Borrowing:  false,
	}
	if err := c.Bind(m); err != nil {
		return nil, err
	}

	// Insert the device in database
	if err := db.DCol.Insert(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func (db *MongoDB) updateDeviceDB(id bson.ObjectId, nm *model.Device) {
	if err := db.DCol.Update(bson.M{"_id": id}, &nm); err != nil {
		return
	}
}

func (db *MongoDB) removeDeviceDB(c echo.Context) {
	id := getID(c)
	m := model.Device{}

	if err := db.DCol.Find(bson.M{"_id": id}).One(&m); err != nil {
		return
	}
	removeFile(&m)

	// Remove the device in database
	if err := db.DCol.RemoveId(id); err != nil {
		return
	}
}

func (db *MongoDB) findDevicesDB() ([]*model.Device, error) {
	m := []*model.Device{}
	// Find all device in database
	if err := db.DCol.Find(nil).Sort("borrowing", "-lastUpdate").All(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func getID(c echo.Context) bson.ObjectId {
	id := bson.ObjectIdHex(c.Param("id"))
	return id
}
