package api

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// RemoveDB is
func (h *Handler) RemoveMacInDB(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	// Remove device in DB
	if err = db.DB("macinodds").C("devices").RemoveId(id); err != nil {
		return
	}

	return c.JSON(http.StatusOK, err)
}
