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

// Login using jwt token of Google.
func (db *MongoDB) Login(c echo.Context) (err error) {
	login := &model.TokenGoogle{}
	if err := c.Bind(login); err != nil {
		return err
	}

	tokenInfo, err := getInfo(login.Token)
	if err != nil {
		return err
	}

	user, firstLogin := db.CheckUser(tokenInfo)

	token, err := genToken(user)
	if err != nil {
		return err
	}

	res := &model.TokenRes{
		Token:      token,
		FirstLogin: firstLogin,
	}

	return c.JSON(http.StatusOK, res)
}

// CheckUser detect the user in database and returns the user's status.
func (db *MongoDB) CheckUser(tokenInfo *oauth2.Tokeninfo) (*model.User, bool) {
	firstLogin := false
	user, status := db.FindUser(tokenInfo)
	if status == "new" {
		// Create a new user to database
		user = db.CreateUser(tokenInfo)
		firstLogin = true
	} else if status == "notComplete" {
		firstLogin = true
	}

	return user, firstLogin
}

// Register the user to database.
func (db *MongoDB) Register(c echo.Context) (err error) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	uid := GetIDFromToken(c)
	db.UpdateUser(uid, u)

	return c.JSON(http.StatusOK, &u)
}

// Logout user and blacklist token
func (db *MongoDB) Logout(c echo.Context) (err error) {
	bl := &model.BlackList{}
	if err := c.Bind(bl); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &bl)
}

// GetIDFromToken return ID of user from jwt token.
func GetIDFromToken(c echo.Context) bson.ObjectId {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uid := claims["id"].(string)
	id := bson.ObjectIdHex(uid)

	return id
}

func verifyAudience(aud string) bool {
	return aud == clientID
}

func getInfo(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(&http.Client{})
	if err != nil {
		return nil, err
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfo, err := tokenInfoCall.IdToken(idToken).Do()
	if err != nil {
		return nil, err
	}

	if !verifyAudience(tokenInfo.Audience) {
		log.Println("token expire!!")
		return nil, nil
	}

	return tokenInfo, nil
}

func genToken(user *model.User) (string, error) {
	// Set custom claims
	claims := &model.JwtCustomClaims{
		user.ID,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(
		[]byte("sMJuczqQPYzocl1s6SLj"),
	)
	if err != nil {
		return "", err
	}

	return t, nil
}
