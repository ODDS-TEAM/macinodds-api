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

// LogIn using jwt token of Google.
func (db *MongoDB) LogIn(c echo.Context) (err error) {
	var login model.TokenGoogle
	if err := c.Bind(&login); err != nil {
		return err
	}

	tokenInfo, err := getInfo(login.Token)
	if err != nil {
		return err
	}

	firstLogin := false
	user := db.FindUser(tokenInfo)
	if user == nil {
		// Create a new user
		user = db.CreateUser(tokenInfo)
		firstLogin = true
	}

	tok, err := genToken(user)
	if err != nil {
		return err
	}

	res := &model.Token{
		Token:      tok,
		FirstLogin: firstLogin,
	}

	return c.JSON(http.StatusOK, res)
}

// Register user to Database
func (db *MongoDB) Register(c echo.Context) (err error) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	q := bson.M{
		"email": u.Email,
	}
	ch := bson.M{
		"$set": bson.M{
			"name":         u.Name,
			"imgProfile":   u.ImgProfile,
			"slackAccount": u.SlackAccount,
			"tel":          u.Tel,
		},
	}

	if err := db.UCol.Update(q, &ch); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &u)
}

// LogOut user and blacklist token
func (db *MongoDB) LogOut(c echo.Context) (err error) {
	return c.JSON(http.StatusCreated, "ok")
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
	claims := &model.JwtCustomClaims{
		user.ID,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok, err := token.SignedString(
		[]byte("sMJuczqQPYzocl1s6SLj"),
	)
	if err != nil {
		return "", err
	}

	return tok, nil
}
