package api

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// RemoveMacInDB forced to remove data from the database directly.
func (h *Handler) RemoveMacInDB(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("mac_odds_team").C("mac").RemoveId(id); err != nil {
		return
	}

	return c.JSON(http.StatusOK, err)
}
