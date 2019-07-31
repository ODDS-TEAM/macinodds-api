package model

type (
	// BlackList keep token when logout and validate.
	BlackList struct {
		Token string `json:"token" bson:"token"`
	}
)
