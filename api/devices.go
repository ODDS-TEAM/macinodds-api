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
	lookBorrowings := bson.M{
		"$lookup": bson.M{
			"from": "borrowings",
			"let":  bson.M{"id": "$_id"},
			"pipeline": []interface{}{
				bson.M{"$match": bson.M{"$expr": bson.M{"$eq": []interface{}{"$device._id", "$$id"}}}},
				bson.M{"$sort": bson.M{"date": -1}},
				bson.M{"$limit": 1},
				bson.M{"$project": bson.M{"borrower": 1, "_id": 1, "returnDate": 1}},
			},
			"as": "borrower",
		},
	}
	addFieldBorower := bson.M{
		"$addFields": bson.M{
			"borrower": bson.M{
				"_id":   bson.M{"$cond": []interface{}{bson.M{"$toBool": "$borrowing"}, bson.M{"$arrayElemAt": []interface{}{"$borrower.borrower._id", 0}}, ""}},
				"name":  bson.M{"$cond": []interface{}{bson.M{"$toBool": "$borrowing"}, bson.M{"$arrayElemAt": []interface{}{"$borrower.borrower.name", 0}}, ""}},
				"slack": bson.M{"$cond": []interface{}{bson.M{"$toBool": "$borrowing"}, bson.M{"$arrayElemAt": []interface{}{"$borrower.borrower.slackAccount", 0}}, ""}},
				"tel":   bson.M{"$cond": []interface{}{bson.M{"$toBool": "$borrowing"}, bson.M{"$arrayElemAt": []interface{}{"$borrower.borrower.tel", 0}}, ""}},
			},
			"returnDate": bson.M{"$cond": []interface{}{bson.M{"$toBool": "$borrowing"}, bson.M{"$arrayElemAt": []interface{}{"$borrower.returnDate", 0}}, ""}},
		},
	}
	projectDeviceInfo := bson.M{
		"$project": bson.M{
			"_id":        1,
			"name":       1,
			"serial":     1,
			"spec":       1,
			"img":        1,
			"location":   1,
			"lastUpdate": 1,
			"borrowing":  1,
			"returnDate": 1,
			"borrower": bson.M{
				"$arrayElemAt": []interface{}{"$borrower", 0},
			},
		},
	}
	projectHideSubBorrowerInfo := bson.M{
		"$project": bson.M{
			"borrower.borrower":   0,
			"borrower.returnDate": 0,
		},
	}
	sortDevice := bson.M{
		"$sort": bson.M{
			"borrowing": 1,
		},
	}

	query := []bson.M{lookBorrowings, addFieldBorower, projectDeviceInfo, projectHideSubBorrowerInfo,sortDevice}
	data := []interface{}{}
	db.DCol.Pipe(query).All(&data)
	return c.JSON(http.StatusOK, &data)
}

func (db *MongoDB) GetDevicesByID(c echo.Context) (err error) {
	id := c.Param("id")
	data := model.Device{}
	db.DCol.FindId(bson.ObjectIdHex(id)).One(&data)
	return c.JSON(http.StatusOK, &data)
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

func (db *MongoDB) findDevicesDB() ([]*model.DeviceList, error) {
	m := []*model.DeviceList{}
	// Find all device in database
	pipe := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "borrowings",
				"localField":   "_id",
				"foreignField": "device._id",
				"as":           "fromBorrowings",
			},
		},
		// {
		// 	"$replaceRoot": bson.M{
		// 		"newRoot": bson.M{
		// 			"devices":        "$$ROOT",
		// 			"fromBorrowings": "$fromBorrowings",
		// 		},
		// 	},
		// },
		// {
		// 	"$project": bson.M{"devices.fromBorrowings": 0},
		// },
	}

	if err := db.DCol.Pipe(pipe).All(&m); err != nil {
		return nil, err
	}
	log.Println(db.DCol.Pipe(pipe).All(&m))
	// if err := db.DCol.Find(nil).Sort("borrowing", "-lastUpdate").All(&m); err != nil {
	// 	return nil, err
	// }

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
