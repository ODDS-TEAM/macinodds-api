package model

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// TokenGoogle hold the jwt token of googleapi
type TokenGoogle struct {
	Token string `json:"token"`
}

// Token that is generated for the requested resource from api.
type Token struct {
	Token      string `json:"token"`
	FirstLogin bool   `json:"firstLogin"`
}

// JwtCustomClaims built in the payload data with ID and role user.
type JwtCustomClaims struct {
	ID   bson.ObjectId `json:"id"`
	Role string        `json:"role"`
	jwt.StandardClaims
}
