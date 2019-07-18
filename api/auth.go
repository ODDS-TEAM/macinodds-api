package api

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	model "gitlab.odds.team/internship/macinodds-api/model"
	"google.golang.org/api/oauth2/v1"
	"gopkg.in/mgo.v2/bson"
)

const clientID = "15378607653-f9lfgsml8th6lf50jfq93v3v2f4vpkpr.apps.googleusercontent.com"

type TokenGoogle struct {
	Token string `json:"token"`
}

func (db *MongoDB) LogIn(c echo.Context) (err error) {
	var login TokenGoogle
	if err := c.Bind(&login); err != nil {
		return err
	}
	// log.Println("token bind", login)

	tokenInfo, err := GetInfo(login.Token)
	// log.Println("token info", tokenInfo)
	if err != nil {
		// log.Println("token info nil!!")
		return err
	}

	firstLogin := false
	user := db.FindUser(tokenInfo)
	if user == nil {
		// first login
		// log.Println("create new user")
		firstLogin = true
		user = db.CreateUser(tokenInfo)
	}

	tok, err := genToken(user)
	if err != nil {
		return err
	}

	res := &Token{
		Token:      tok,
		FirstLogin: firstLogin,
	}

	return c.JSON(http.StatusOK, res)
}

type Token struct {
	Token      string `json:"token"`
	FirstLogin bool   `json:"firstLogin"`
}

type JwtCustomClaims struct {
	ID   bson.ObjectId `json:"id"`
	Role string        `json:"role"`
	jwt.StandardClaims
}

func genToken(user *model.User) (string, error) {
	claims := &JwtCustomClaims{
		user.ID,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok, err := token.SignedString([]byte("sMJuczqQPYzocl1s6SLj"))
	if err != nil {
		return "", err
	}
	return tok, nil
}

func (db *MongoDB) CreateUser(tokeninfo *oauth2.Tokeninfo) *model.User {
	u := &model.User{
		ID:       bson.NewObjectId(),
		Role:     "indevidual",
		Email:    tokeninfo.Email,
		CreateAt: time.Now(),
	}
	if err := db.UCol.Insert(&u); err != nil {
		return nil
	}

	return u
}

func (db *MongoDB) FindUser(tokeninfo *oauth2.Tokeninfo) *model.User {
	u := &model.User{}

	q := bson.M{
		"email": tokeninfo.Email,
	}

	if err := db.UCol.Find(q).One(&u); err != nil {
		return nil
	}

	return u
}

func verifyAudience(aud string) bool {
	return aud == clientID
}

func GetInfo(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(&http.Client{})
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	// log.Println("tokenInfoCall:", tokenInfoCall)
	tokenInfo, err := tokenInfoCall.IdToken(idToken).Do()

	// log.Println("tokenInfo:", tokenInfo.VerifiedEmail)

	if err != nil {
		return nil, err
	}

	if !verifyAudience(tokenInfo.Audience) {
		log.Println("token expire!!")
		return nil, nil
	}
	return tokenInfo, nil
}

func (db *MongoDB) Register(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}

func (db *MongoDB) LogOut(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
}
