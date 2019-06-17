package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	device "gitlab.odds.team/internship/macinodds-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var devices []device.Device

var mongoHost string = "139.5.146.213:27017"

// Handler
func GetAPI(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Macinodds API v.1.0!")
}

func GetDevices(c echo.Context) (err error) {
	session, err := mgo.Dial(mongoHost)
	defer session.Close()
	s := session.DB("macinodds").C("devices")

	err = s.Find(bson.M{}).All(&devices)
	if err != nil {
		panic(err)
	}

	fmt.Println(devices)
	return c.JSON(http.StatusOK, devices)
}
