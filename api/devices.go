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
	file, src, _ := openFile(c)
	imgName, filePath := genImgID(file.Filename)
	createFile(filePath, src)

	m, err := db.insertDeviceDB(c, imgName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

// UpdateDevice updates the entry for a given device.
func (db *MongoDB) UpdateDevice(c echo.Context) (err error) {
	id, newm, m := db.compareDevices(c)

	if newm.Img != m.Img || newm.Img == "" {
		file, src, _ := openFile(c)
		imgNewName, filePath := genImgID(file.Filename)
		createFile(filePath, src)

		newm.Img = imgNewName
		if m.Img != "" {
			removeFile(m)
		}
	}
	db.updateDeviceDB(id, newm)

	return c.JSON(http.StatusOK, &newm)
}

// RemoveDevice removes a given device by its ID.
func (db *MongoDB) RemoveDevice(c echo.Context) (err error) {
	m := db.removeDeviceDB(c)
	removeFile(m)

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

func (db *MongoDB) GetDevicesByID(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) BorrowDevice(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) ReturnDevice(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) insertDeviceDB(c echo.Context, i string) (*model.Device, error) {
	m := &model.Device{
		ID:         bson.NewObjectId(),
		Img:        i,
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
	// Update Mac in database
	if err := db.DCol.UpdateId(id, &nm); err != nil {
		return
	}
}

func (db *MongoDB) removeDeviceDB(c echo.Context) *model.Device {
	id := getID(c)
	m := model.Device{}

	if err := db.DCol.Find(bson.M{"_id": id}).One(&m); err != nil {
		return nil
	}

	// Remove the device in database
	if err := db.DCol.RemoveId(id); err != nil {
		return nil
	}

	return &m
}

func (db *MongoDB) findDevicesDB() ([]*model.Device, error) {
	m := []*model.Device{}
	// Find all device in database
	if err := db.DCol.Find(nil).Sort("borrowing", "-lastUpdate").All(&m); err != nil {
		return nil, err
	}

	return m, nil
}

// find1

func (db *MongoDB) compareDevices(c echo.Context) (bson.ObjectId, *model.Device, *model.Device) {
	id := getID(c)
	newm := &model.Device{
		ID:         id,
		LastUpdate: time.Now(),
	}
	if err := c.Bind(newm); err != nil {
		return id, nil, nil
	}

	m := &model.Device{}
	if err := db.DCol.FindId(id).One(&m); err != nil {
		return id, newm, nil
	}

	log.Println("new image:", newm.Img, "| old image:", m.Img)

	return id, newm, m
}

func getID(c echo.Context) bson.ObjectId {
	i := c.Param("id")
	if i == "" {
		return ""
	}

	id := bson.ObjectIdHex(c.Param("id"))

	return id
}
